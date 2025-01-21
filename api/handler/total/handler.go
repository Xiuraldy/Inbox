package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "challenge/usecase/total"
)

type TotalHandler struct {
	usecase usecase.UseCase
}

func NewTotal(usecase usecase.UseCase) *TotalHandler {
	return &TotalHandler{
		usecase: usecase,
	}
}

func (e *TotalHandler) GetTotal(w http.ResponseWriter, r *http.Request) {
	fromStr := r.URL.Query().Get("from")
	from, err := strconv.Atoi(fromStr)
	if err != nil || from < 0 {
		http.Error(w, "Invalid from value", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search")

	totalResponse, err := e.usecase.GetTotal(from, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(totalResponse)
}
