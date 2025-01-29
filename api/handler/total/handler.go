package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "challenge/usecase/total"
)

// Manejar las solicitudes HTTP relacionadas con el número total de emails.
type TotalHandler struct {
	usecase usecase.UseCase
}

// Crea y devuelve una nueva instancia de TotalHandler con el usecase inyectado
func NewTotal(usecase usecase.UseCase) *TotalHandler {
	return &TotalHandler{
		usecase: usecase,
	}
}

// Maneja las solicitudes HTTP para obtener el número total de emails con filtros.
func (e *TotalHandler) GetTotal(w http.ResponseWriter, r *http.Request) {
	fromStr := r.URL.Query().Get("from") // Obtiene el parámetro "from"
	from, err := strconv.Atoi(fromStr)   // Convierte el parámetro "from" de string a entero
	if err != nil || from < 0 {
		http.Error(w, "Invalid from value", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search") // Obtiene el parámetro "search"

	// Llama al caso de uso para obtener el número total de emails con los filtros aplicados
	totalResponse, err := e.usecase.GetTotal(from, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Configura el encabezado de la respuesta como JSON
	json.NewEncoder(w).Encode(totalResponse)           // Devuelve la respuesta codificada en JSON
}
