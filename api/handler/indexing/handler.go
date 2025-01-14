package handler

import (
	usecase "challengeTruora/usecase/indexing"
	"encoding/json"
	"net/http"
)

type IndexingHandler struct {
	usecase usecase.UseCase
}

func NewIndexing(usecase usecase.UseCase) *IndexingHandler {
	return &IndexingHandler{
		usecase: usecase,
	}
}

func (e *IndexingHandler) Indexing(w http.ResponseWriter, r *http.Request) {
	err := e.usecase.Indexing()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
