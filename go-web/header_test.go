package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	w.Header().Add("X-Powered-By", "Ibnnu Dev")
	fmt.Fprint(w, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Go-http-client/1.1")
	request.Header.Add("X-Powered-By", "Ibnnu Dev")
	response := httptest.NewRecorder()

	Handler(response, request)

	fmt.Println("X-Powered-By:", response.Header().Get("X-Powered-By"))
	assert.Equal(t, "Ibnnu Dev", response.Header().Get("X-Powered-By"))
}

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	fmt.Fprintf(w, "My name is %s %s", firname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=Ibnnu&lastname=Dev")
	request := httptest.NewRequest(http.MethodPost, "/hello", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	FormPost(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println("Response Body:", string(body))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "My name is Ibnnu Dev", string(body))
	responseBody := string(body)
	assert.Contains(t, responseBody, "Ibnnu")
	assert.Contains(t, responseBody, "Dev")
}
