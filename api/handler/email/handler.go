package handler

import (
	usecase "challenge/usecase/email"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type EmailHandler struct {
	usecase usecase.UseCase
}

func NewEmail(usecase usecase.UseCase) *EmailHandler {
	return &EmailHandler{
		usecase: usecase,
	}
}

func (e *EmailHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	emailResponse, err := e.usecase.GetEmail()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emails, ok := emailResponse["data"].([]interface{})
	if !ok {
		http.Error(w, "Invalid data format", http.StatusInternalServerError)
		return
	}

	totalCount, ok := emailResponse["totalCount"].(int)
	if !ok {
		http.Error(w, "Invalid totalCount format", http.StatusInternalServerError)
		return
	}

	search := r.URL.Query().Get("search")
	if search != "" {
		emails = searchEmail(emails, search)
		totalCount = len(emails)
	}

	paginator, err := strconv.Atoi(r.URL.Query().Get("paginator"))
	if err != nil {
		http.Error(w, "Invalid paginator value", http.StatusBadRequest)
		return
	}

	finishPag := paginator
	if finishPag > len(emails) {
		finishPag = len(emails)
	}
	initialPag := paginator - 5

	response := map[string]interface{}{
		"emails":     emails[initialPag:finishPag],
		"totalCount": totalCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func searchEmail(emails []interface{}, search string) []interface{} {
	var results []interface{}
	for _, email := range emails {
		if mapEmail, ok := email.(map[string]interface{}); ok {
			for _, value := range mapEmail {
				if strings.Contains(fmt.Sprintf("%v", value), search) {
					results = append(results, email)
					break
				}
			}
		}
	}
	return results
}
