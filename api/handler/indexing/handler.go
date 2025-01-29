package handler

import (
	usecase "challenge/usecase/indexing"
	"encoding/json"
	"net/http"
)

// Manejar las solicitudes HTTP relacionadas con la indexación.
type IndexingHandler struct {
	usecase usecase.UseCase // Inyecta el caso de uso de indexación
}

// Crea y devuelve una nueva instancia de IndexingHandler con el usecase inyectado
func NewIndexing(usecase usecase.UseCase) *IndexingHandler {
	return &IndexingHandler{
		usecase: usecase,
	}
}

// Maneja las solicitudes HTTP para iniciar el proceso de indexación.
func (e *IndexingHandler) Indexing(w http.ResponseWriter, r *http.Request) {
	// Ejecuta la lógica de indexación a través del usecase
	err := e.usecase.Indexing()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Configura el encabezado de la respuesta como JSON

	json.NewEncoder(w).Encode(nil) // Configura el encabezado de la respuesta como JSON
}
