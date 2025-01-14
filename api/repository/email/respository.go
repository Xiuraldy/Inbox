package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Repository interface {
	GetEmail() ([]interface{}, error)
}

type repositoryEmail struct {
	username string
	password string
	org      string
	stream   string
	url      string
}

func NewRepositoryEmail(username, password, org, stream string) Repository {
	url := fmt.Sprintf("https://api.openobserve.ai/api/%s/_search", org)
	return &repositoryEmail{
		username: username,
		password: password,
		org:      org,
		stream:   stream,
		url:      url,
	}
}

func (r *repositoryEmail) GetEmail() ([]interface{}, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":        fmt.Sprintf("SELECT * FROM %s", r.stream),
			"start_time": 1674789786006000,
			"end_time":   1674789786006000,
			"from":       0,
			"size":       10,
		},
		"search_type": "ui",
		"timeout":     0,
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error codificando consulta: %w", err)
	}

	req, err := http.NewRequest("POST", r.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %w", err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", r.username, r.password)))
	req.Header.Add("Authorization", "Bearer "+auth)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error enviando solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("estado inesperado: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta: %w", err)
	}

	log.Printf("Respuesta recibida: %s", body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parseando respuesta: %w", err)
	}

	hits, ok := result["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("respuesta inesperada: 'hits' no encontrado o tipo incorrecto")
	}

	return hits, nil
}
