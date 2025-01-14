package main

import (
	"challenge/shared"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	shared.SetupRouter(r, os.Getenv("EMAIL"), os.Getenv("PASSWORD"), os.Getenv("ORG"), "default")

	http.ListenAndServe(":8080", r)

	log.Println("Procesamiento completado.")
}
