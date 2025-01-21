package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "challenge/usecase/email"
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
	fromStr := r.URL.Query().Get("from")
	from, err := strconv.Atoi(fromStr)
	if err != nil || from < 0 {
		http.Error(w, "Invalid from value", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search")

	emailResponse, err := e.usecase.GetEmail(from, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emailResponse)
}
