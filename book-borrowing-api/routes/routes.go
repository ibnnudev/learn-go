package routes

import (
	"learn-go/book-borrowing-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, h *handlers.BookHandler) {
	r.HandleFunc("/books", h.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/books", h.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", h.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")
}
