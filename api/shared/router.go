package shared

import (
	handlerEmail "challenge/handler/email"
	handlerIndexing "challenge/handler/indexing"
	handlerTotal "challenge/handler/total"
	repositoryEmail "challenge/repository/email"
	repositoryTotal "challenge/repository/total"
	usecaseEmail "challenge/usecase/email"
	usecaseIndexing "challenge/usecase/indexing"
	usecaseTotal "challenge/usecase/total"

	"github.com/go-chi/chi"
)

// Configura el router y las dependencias
func SetupRouter(r *chi.Mux, username, password, stream, url string) {
	emailHandler := initEmailModule(username, password, stream, url)
	totalHandler := initTotalModule(username, password, stream, url)
	indexingHandler := initIndexingModule()

	r.Use(Cors())
	r.Get("/email", emailHandler.GetEmail)
	r.Get("/total", totalHandler.GetTotal)
	r.Get("/indexing", indexingHandler.Indexing)
}

// Modularización de inicialización de dependencias
func initEmailModule(username, password, stream, url string) *handlerEmail.EmailHandler {
	repo := repositoryEmail.NewRepositoryEmail(username, password, stream, url)
	usecase := usecaseEmail.NewUseCaseEmail(repo)
	return handlerEmail.NewEmail(usecase)
}

func initTotalModule(username, password, stream, url string) *handlerTotal.TotalHandler {
	repo := repositoryTotal.NewRepositoryTotal(username, password, stream, url)
	usecase := usecaseTotal.NewUseCaseTotal(repo)
	return handlerTotal.NewTotal(usecase)
}

func initIndexingModule() *handlerIndexing.IndexingHandler {
	usecase := usecaseIndexing.NewUseCaseIndexing()
	return handlerIndexing.NewIndexing(usecase)
}
