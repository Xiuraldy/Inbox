package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repository interface {
	GetEmail(from int, search string) ([]interface{}, error)
}

type repositoryEmail struct {
	username string
	password string
	stream   string
	url      string
}

func NewRepositoryEmail(username, password, stream, url string) Repository {
	return &repositoryEmail{
		username: username,
		password: password,
		stream:   stream,
		url:      url,
	}
}

func (r *repositoryEmail) GetEmail(from int, search string) ([]interface{}, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM %s", r.stream)
	if search != "" {
		sqlQuery += fmt.Sprintf(
			" WHERE body LIKE '%%%s%%' OR subject LIKE '%%%s%%' OR from LIKE '%%%s%%' OR date LIKE '%%%s%%' OR message_id LIKE '%%%s%%' OR to LIKE '%%%s%%' OR _timestamp::text LIKE '%%%s%%'",
			search, search, search, search, search, search, search,
		)

	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":        sqlQuery,
			"start_time": 1737060892521866,
			"end_time":   1737152148243135,
			"from":       from,
			"size":       5,
		},
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	req, err := http.NewRequest("POST", r.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	credentials := fmt.Sprintf("%s:%s", r.username, r.password)
	credentialsBase64 := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", credentialsBase64))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status: %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	type Response struct {
		Hits []interface{} `json:"hits"`
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return response.Hits, nil
}
