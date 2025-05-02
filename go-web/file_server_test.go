package goweb

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("resources")
	fs := http.FileServer(directory)

	mux := http.NewServeMux()

	mux.Handle("/resources/", http.StripPrefix("/resources/", fs))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}
}
