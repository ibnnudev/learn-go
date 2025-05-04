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

func UserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("id")
	fmt.Fprintf(w, "Halo, id saya:%s", userId)
}

func ProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	price := p.ByName("price")
	fmt.Fprintln(w, "PRODUCT")
	fmt.Fprintln(w, "Name:", name)
	fmt.Fprintln(w, "Price:", price)

	fmt.Fprintln(w, "===================================")
}

func TestUserHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:id", UserHandler)

	request := httptest.NewRequest("GET", "/user/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Halo, id saya:123", string(body))
}

func TestProductHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:name/:price", ProductHandler)

	request := httptest.NewRequest("GET", "/product/jeruk/10000", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "PRODUCT\nName: jeruk\nPrice: 10000\n===================================\n", string(body))
}
