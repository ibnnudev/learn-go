package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Panic occurred:%s", i)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("This is a panic")
	})

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Panic occurred:This is a panic", string(body))
}

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page not found")
	})

	request := httptest.NewRequest("GET", "/nonexistent", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Page not found", string(body))
}

func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		fmt.Println("Method not allowed")
		fmt.Println("Request method:", r.Method)
		fmt.Println("Request URL:", r.URL)
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello, HTTP Router!")
	})

	request := httptest.NewRequest("POST", "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Method not allowed", string(body))
}
