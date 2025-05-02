package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "main page")
	})
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about page")
	})
	mux.HandleFunc("/about/employee", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about > employee page")
	})

	mux.HandleFunc("/about/employee/:id", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/about/employee/"):]
		fmt.Fprintf(w, "about > employee page %s", id)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestReques(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
