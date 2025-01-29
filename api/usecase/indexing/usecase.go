package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

// Estructura del Email
type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

// Pool de buffers para reutilizar memoria al codificar JSON, optimizando rendimiento
var jsonBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Interfaz que define la funcionalidad del caso de uso
type UseCase interface {
	Indexing() error // Método para la indexación de correos
}

// Implementación del caso de uso para la indexación
type useCaseIndexing struct {
}

// Constructor del caso de uso, retorna una nueva instancia
func NewUseCaseIndexing() UseCase {
	return &useCaseIndexing{}
}

// Inicia la indexación de emails
func (u *useCaseIndexing) Indexing() error {
	// Inicia servidor de profiling
	go func() {
		log.Println("pprof server iniciado en http://localhost:6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	root := os.Getenv("ROOT")

	// Canal para manejar rutas de archivos en paralelo
	emailPaths := make(chan string, 100)
	var wg sync.WaitGroup

	// Crea 4 goroutines (hilos) que procesarán los archivos concurrentemente
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range emailPaths {
				processEmail(path) // Procesa cada archivo de email
			}
		}()
	}

	// Recorre todos los archivos del directorio raíz y los envía al canal
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			emailPaths <- path
		}
		return nil
	})
	close(emailPaths) // Cierra el canal después de recorrer todos los archivos

	if err != nil {
		log.Fatalf("Error while browsing the files: %v", err)
	}

	wg.Wait() // Espera a que todas las goroutines finalicen

	return nil
}

// Procesa un email desde su ruta de archivo
func processEmail(path string) {
	// Lee el contenido del archivo
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Warning: The file cannot be read %s: %v", path, err)
		return
	}

	// Parsea el contenido del archivo y lo convierte en una estructura `Email`
	email, err := parseEmail(string(content))
	if err != nil {
		log.Printf("Warning: Error parsing email %s: %v", path, err)
		return
	}

	// Envía el email a indexarse en el servidor
	err = indexEmail(email)
	if err != nil {
		log.Printf("Warning: Error indexing email %s: %v", path, err)
	}
}

// Parsea el contenido del email separando los encabezados del cuerpo
func parseEmail(content string) (*Email, error) {
	// Divide el contenido del email en encabezado y cuerpo, basándose en líneas en blanco
	re := regexp.MustCompile(`\r?\n\r?\n+`)
	parts := re.Split(content, 2)

	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid formatting: No headers and body found")
	}

	headers := parts[0]
	body := parts[1]

	email := &Email{}
	// Itera sobre cada línea del encabezado para extraer los valores y asignarlos
	for _, line := range strings.Split(headers, "\n") {
		switch {
		case strings.HasPrefix(line, "Message-ID:"):
			email.MessageID = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID:"))
		case strings.HasPrefix(line, "Date:"):
			email.Date = strings.TrimSpace(strings.TrimPrefix(line, "Date:"))
		case strings.HasPrefix(line, "From:"):
			email.From = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
		case strings.HasPrefix(line, "To:"):
			email.To = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
		case strings.HasPrefix(line, "Subject:"):
			email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
		}
	}
	email.Body = body // Guarda el cuerpo del email
	return email, nil
}

// Envia el email a un servidor para su indexación
func indexEmail(email *Email) error {
	// Carga variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtiene un buffer del pool para optimizar la conversión a JSON
	buffer := jsonBufferPool.Get().(*bytes.Buffer)
	defer jsonBufferPool.Put(buffer) // Devuelve el buffer al pool para su reutilización
	buffer.Reset()                   // Limpia el buffer

	// Convierte el email a formato JSON
	err = json.NewEncoder(buffer).Encode(email)
	if err != nil {
		return fmt.Errorf("error converting to JSON: %v", err)
	}

	// Crea la solicitud HTTP para enviar el email al servidor
	req, err := http.NewRequest("POST", os.Getenv("URL"), buffer)
	if err != nil {
		return fmt.Errorf("error creating the request: %v", err)
	}

	// Añade credenciales de autenticación
	req.SetBasicAuth(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	// Envía la solicitud HTTP al servidor
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending the request: %v", err)
	}
	defer resp.Body.Close() // Asegura el cierre del cuerpo de la respuesta

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error in the answer: %s", string(body))
	}

	log.Printf("Mail indexed correctly: %s", email.MessageID) // Mensaje de éxito
	return nil
}
