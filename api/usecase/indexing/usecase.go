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

type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

var jsonBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type UseCase interface {
	Indexing() error
}

type useCaseIndexing struct {
}

func NewUseCaseIndexing() UseCase {
	return &useCaseIndexing{}
}

func (u *useCaseIndexing) Indexing() error {
	go func() {
		log.Println("pprof server iniciado en http://localhost:6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	root := os.Getenv("ROOT")

	emailPaths := make(chan string, 100)
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range emailPaths {
				processEmail(path)
			}
		}()
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			emailPaths <- path
		}
		return nil
	})
	close(emailPaths)

	if err != nil {
		log.Fatalf("Error recorriendo los archivos: %v", err)
	}

	wg.Wait()

	return nil
}

func processEmail(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Advertencia: No se puede leer el archivo %s: %v", path, err)
		return
	}

	email, err := parseEmail(string(content))
	if err != nil {
		log.Printf("Advertencia: Error parseando el email %s: %v", path, err)
		return
	}

	err = indexEmail(email)
	if err != nil {
		log.Printf("Advertencia: Error indexando el email %s: %v", path, err)
	}
}

func parseEmail(content string) (*Email, error) {
	re := regexp.MustCompile(`\r?\n\r?\n+`)
	parts := re.Split(content, 2)

	if len(parts) < 2 {
		return nil, fmt.Errorf("formato inválido: No se encontraron encabezados y cuerpo")
	}

	headers := parts[0]
	body := parts[1]

	email := &Email{}
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
	email.Body = body
	return email, nil
}

func indexEmail(email *Email) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	buffer := jsonBufferPool.Get().(*bytes.Buffer)
	defer jsonBufferPool.Put(buffer)
	buffer.Reset()

	err = json.NewEncoder(buffer).Encode(email)
	if err != nil {
		return fmt.Errorf("error convirtiendo a JSON: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("URL"), buffer)
	if err != nil {
		return fmt.Errorf("error creando la solicitud: %v", err)
	}

	req.SetBasicAuth(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error enviando la solicitud: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	log.Printf("Correo indexado correctamente: %s", email.MessageID)
	return nil
}
