package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "challenge/usecase/email"
)

// Manejar las solicitudes HTTP relacionadas con emails.
type EmailHandler struct {
	usecase usecase.UseCase
}

// Crea y devuelve una nueva instancia de EmailHandler con el usecase inyectado
func NewEmail(usecase usecase.UseCase) *EmailHandler {
	return &EmailHandler{
		usecase: usecase,
	}
}

// Maneja las solicitudes HTTP para obtener emails con filtros.
func (e *EmailHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	fromStr := r.URL.Query().Get("from") // Obtiene el parámetro "from"
	from, err := strconv.Atoi(fromStr)   // Convierte el parámetro "from" de string a entero
	if err != nil || from < 0 {
		http.Error(w, "Invalid from value", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search") // Obtiene el parámetro "search"

	// Llama al caso de uso para obtener los emails con los filtros aplicados
	emailResponse, err := e.usecase.GetEmail(from, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Configura el encabezado de la respuesta como JSON
	json.NewEncoder(w).Encode(emailResponse)           // Devuelve la respuesta codificada en JSON
}
