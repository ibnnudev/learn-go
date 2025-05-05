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

type LogMiddleware struct {
	Next http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL)
	middleware.Next.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello, World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	middleware := LogMiddleware{
		Next: router,
	}

	middleware.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello, World!", string(body))
}
