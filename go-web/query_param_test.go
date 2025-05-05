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

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "Hello, %s!\n", name)
	fmt.Fprintf(w, "You are %s years old.\n", age)
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=John&age=30", nil)
	response := httptest.NewRecorder()

	SayHello(response, request)
	result := response.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))

	assert.Equal(t, "Hello, John!\nYou are 30 years old.\n", string(body))

}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var names = query["name"]
	fmt.Fprint(w, strings.Join(names, ", "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=John&name=Jane", nil)
	response := httptest.NewRecorder()

	MultipleParameterValues(response, request)
	result := response.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))

	assert.Equal(t, "John, Jane", string(body))
}
