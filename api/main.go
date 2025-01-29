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
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Crear router
	r := chi.NewRouter()

	// Declarar variables de entorno
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	stream := os.Getenv("STREAM")
	search := os.Getenv("SEARCH")

	// Configurar rutas
	shared.SetupRouter(r, email, password, stream, search)

	// Levantar servidor
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error when starting the server:", err)
	}
}
