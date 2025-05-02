package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "File path is required")
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filePath))
	http.ServeFile(w, r, "./resources/"+filePath)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer server.Close()
}
