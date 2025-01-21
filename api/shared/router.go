package shared

import (
	handler "challenge/handler/email"
	handlerIndexing "challenge/handler/indexing"
	handlerTotal "challenge/handler/total"
	repository "challenge/repository/email"
	repositoryTotal "challenge/repository/total"
	usecase "challenge/usecase/email"
	usecaseIndexing "challenge/usecase/indexing"
	usecaseTotal "challenge/usecase/total"

	"github.com/go-chi/chi"
)

func SetupRouter(r *chi.Mux, username, password, stream, url string) {
	repositoryEmail := repository.NewRepositoryEmail(username, password, stream, url)
	usecaseEmail := usecase.NewUseCaseEmail(repositoryEmail)
	handlerEmail := handler.NewEmail(usecaseEmail)

	repositoryTotal := repositoryTotal.NewRepositoryTotal(username, password, stream, url)
	usecaseTotal := usecaseTotal.NewUseCaseTotal(repositoryTotal)
	handlerTotal := handlerTotal.NewTotal(usecaseTotal)

	usecaseIndexing := usecaseIndexing.NewUseCaseIndexing()
	handlerIndexing := handlerIndexing.NewIndexing(usecaseIndexing)

	r.Use(Cors())
	r.Get("/email", handlerEmail.GetEmail)
	r.Get("/total", handlerTotal.GetTotal)
	r.Get("/indexing", handlerIndexing.Indexing)
}
