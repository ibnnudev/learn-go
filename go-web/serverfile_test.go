package goweb

import (
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
