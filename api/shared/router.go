package shared

import (
	handler "challengeTruora/handler/email"
	handlerIndexing "challengeTruora/handler/indexing"
	repository "challengeTruora/repository/email"
	usecase "challengeTruora/usecase/email"
	usecaseIndexing "challengeTruora/usecase/indexing"

	"github.com/go-chi/chi"
)

func SetupRouter(r *chi.Mux, username, password, org, stream string) {
	repositoryEmail := repository.NewRepositoryEmail(username, password, org, stream)
	usecase := usecase.NewUseCaseEmail(repositoryEmail)
	handlerEmail := handler.NewEmail(usecase)

	indexingUsecase := usecaseIndexing.NewUseCaseIndexing()
	indexingHandler := handlerIndexing.NewIndexing(indexingUsecase)

	r.Use(Cors())
	r.Get("/email", handlerEmail.GetEmail)
	r.Get("/indexing", indexingHandler.Indexing)
}
