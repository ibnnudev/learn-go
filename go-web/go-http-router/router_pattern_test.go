package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Home Page")
}

func BookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "Book Page")
}

func DetailBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookName := ps.ByName("name")
	fmt.Fprintln(w, "Book Name:", bookName)
}

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/", HomeHandler)
	router.GET("/book", BookHandler)
	router.GET("/book/:name", DetailBookHandler)

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	if string(body) != "Home Page\n" {
		t.Errorf("Expected 'Home Page', got '%s'", string(body))
	}

	request = httptest.NewRequest("GET", "/book", nil)
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response = recorder.Result()

	body, _ = io.ReadAll(response.Body)
	if string(body) != "Book Page\n" {
		t.Errorf("Expected 'Book Page', got '%s'", string(body))
	}

	request = httptest.NewRequest("GET", "/book/golang", nil)
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response = recorder.Result()

	body, _ = io.ReadAll(response.Body)
	if string(body) != "Book Name: golang\n" {
		t.Errorf("Expected 'Book Name: golang', got '%s'", string(body))
	}
}
