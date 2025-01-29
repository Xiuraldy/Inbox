// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof" // Paquete para profiling
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strings"

// 	"github.com/joho/godotenv"
// )

// type Email struct {
// 	MessageID               string `json:"message_id"`
// 	Date                    string `json:"date"`
// 	From                    string `json:"from"`
// 	To                      string `json:"to"`
// 	Subject                 string `json:"subject"`
// 	MimeVersion             string `json:"mime_version"`
// 	ContentType             string `json:"content_type"`
// 	ContentTransferEncoding string `json:"content_transfer_encoding"`
// 	XFrom                   string `json:"x_from"`
// 	XTo                     string `json:"x_to"`
// 	Xcc                     string `json:"x_cc"`
// 	Xbcc                    string `json:"x_bcc"`
// 	XFolder                 string `json:"x_folder"`
// 	XOrigin                 string `json:"x_origin"`
// 	XFileName               string `json:"x_filename"`
// 	Body                    string `json:"body"`
// }

// func main() {
// 	go func() {
// 		log.Println("pprof server iniciado en http://localhost:6060")
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()

// 	root := "//?/C:/Users/Xiuraldy/Documents/Programación/challenge/api/data/enron_mail_20110402/maildir"

// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if !info.IsDir() {
// 			processEmail(path)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatalf("Error recorriendo los archivos: %v", err)
// 	}

// 	log.Println("Procesamiento completado.")

// 	fmt.Println("Presiona Enter para salir...")
// 	_, _ = fmt.Scanln()
// }

// func processEmail(path string) {
// 	content, err := os.ReadFile(path)
// 	if err != nil {
// 		log.Printf("Advertencia: No se puede leer el archivo %s: %v", path, err)
// 		return
// 	}

// 	email, err := parseEmail(string(content))
// 	if err != nil {
// 		log.Printf("Advertencia: Error parseando el email %s: %v", path, err)
// 		return
// 	}

// 	err = indexEmail(email)
// 	if err != nil {
// 		log.Printf("Advertencia: Error indexando el email %s: %v", path, err)
// 	}
// }

// func parseEmail(content string) (*Email, error) {
// 	re := regexp.MustCompile(`\r?\n\r?\n+`)
// 	parts := re.Split(content, 2)

// 	if len(parts) < 2 {
// 		return nil, fmt.Errorf("Formato inválido: No se encontraron encabezados y cuerpo")
// 	}

// 	headers := parts[0]
// 	body := parts[1]

// 	email := &Email{}
// 	for _, line := range strings.Split(headers, "\n") {
// 		if strings.HasPrefix(line, "Message-ID:") {
// 			email.MessageID = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID:"))
// 		} else if strings.HasPrefix(line, "Date:") {
// 			email.Date = strings.TrimSpace(strings.TrimPrefix(line, "Date:"))
// 		} else if strings.HasPrefix(line, "From:") {
// 			email.From = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
// 		} else if strings.HasPrefix(line, "To:") {
// 			email.To = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
// 		} else if strings.HasPrefix(line, "Subject:") {
// 			email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
// 		}
// 	}

// 	email.Body = body
// 	return email, nil
// }

// func indexEmail(email *Email) error {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	jsonData, err := json.Marshal(email)
// 	if err != nil {
// 		return fmt.Errorf("Error convirtiendo a JSON: %v", err)
// 	}

// 	req, err := http.NewRequest("POST", os.Getenv("URL"), bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("Error creando la solicitud: %v", err)
// 	}

// 	req.SetBasicAuth(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("Error enviando la solicitud: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
// 		body, _ := io.ReadAll(resp.Body)
// 		return fmt.Errorf("Error en la respuesta: %s", string(body))
// 	}

// 	log.Printf("Correo indexado correctamente: %s", email.MessageID)
// 	return nil
// }