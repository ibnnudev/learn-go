package main

import (
	"learn-go/book-borrowing-api/database"
	"learn-go/book-borrowing-api/handlers"
	"learn-go/book-borrowing-api/models"
	"learn-go/book-borrowing-api/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.ConnectDB()

	database.MigrateDB(&models.Book{})

	r := mux.NewRouter()

	bookHandler := &handlers.BookHandler{DB: database.DB}

	routes.SetupRoutes(r, bookHandler)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf(("Server started on port %s\n"), port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
