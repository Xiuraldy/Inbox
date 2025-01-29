package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define la interfaz para acceder al total de emails
type Repository interface {
	GetTotal(from int, search string) ([]interface{}, error) // Realiza una consulta en la base de datos con filtros
}

// Implementación del repositorio para el total de emails
type repositoryTotal struct {
	username string
	password string
	stream   string
	url      string
	client   *http.Client
}

// Crea una nueva instancia de repositoryTotal
func NewRepositoryTotal(username, password, stream, url string) Repository {
	return &repositoryTotal{
		username: username,
		password: password,
		stream:   stream,
		url:      url,
		client:   &http.Client{},
	}
}

// Realiza una consulta SQL en la API y devuelve el total de emails
func (r *repositoryTotal) GetTotal(from int, search string) ([]interface{}, error) {
	sqlQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.stream) // Crea la consulta SQL base
	// Si se proporciona un término de búsqueda las añade a la consulta
	if search != "" {
		sqlQuery += fmt.Sprintf(
			" WHERE body LIKE '%%%s%%' OR subject LIKE '%%%s%%' OR from LIKE '%%%s%%' OR date LIKE '%%%s%%' OR message_id LIKE '%%%s%%' OR to LIKE '%%%s%%' OR _timestamp::text LIKE '%%%s%%'",
			search, search, search, search, search, search, search,
		)
	}

	// Construir el payload para la consulta
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":        sqlQuery,
			"start_time": 1737060892521866,
			"end_time":   1737152148243135,
			"from":       from,
			"size":       5,
		},
	}

	// Convertir la consulta a formato JSON
	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	// Crear el request HTTP con el payload JSON
	req, err := http.NewRequest("POST", r.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Configurar encabezados para autenticación y tipo de contenido
	credentials := fmt.Sprintf("%s:%s", r.username, r.password)
	credentialsBase64 := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", credentialsBase64))

	// Enviar la solicitud al servidor
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Verificar si la respuesta HTTP tiene un código válido
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status: %d, response: %s", resp.StatusCode, string(body))
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Estructura para deserializar la respuesta
	type Response struct {
		Hits []interface{} `json:"hits"`
	}

	var response Response
	// Convertir el cuerpo JSON de la respuesta en una estructura Go
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return response.Hits, nil
}
