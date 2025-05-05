package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("file") != "" && r.URL.Query().Get("file") == "ok" {
		http.ServeFile(w, r, "./resources/ok.txt")
	} else {
		http.Error(w, "File not found", http.StatusNotFound)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Close()
}

//go:embed resources/ok.txt
var okFile string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("file") != "" && r.URL.Query().Get("file") == "ok" {
		fmt.Fprint(w, okFile)
	} else {
		http.Error(w, "File not found", http.StatusNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Close()
}
