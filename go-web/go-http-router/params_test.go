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

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/user/:id", UserHandler)
	request := httptest.NewRequest("GET", "/user/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(nil, "Halo, id saya:123", string(body))
	fmt.Println(string(body))
}
