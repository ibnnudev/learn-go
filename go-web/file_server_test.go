package goweb

import (
	"embed"
	"io/fs"
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

//go:embed resources
var resources embed.FS

func TestEmbedFile(t *testing.T) {
	subFS, _ := fs.Sub(resources, "resources")
	directory := http.FileServer(http.FS(subFS))

	mux := http.NewServeMux()

	mux.Handle("/resources/", http.StripPrefix("/resources/", directory))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}
}
