package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{name} is required")
	} else {
		fmt.Fprint(w, "Hello, ", name)
		w.WriteHeader(http.StatusOK)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "/?name=ibnu", nil)
	response := httptest.NewRecorder()

	ResponseCode(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println("Status Code:", result.Status)
	fmt.Println("Status Code:", result.StatusCode)
	fmt.Println("Response Body:", string(body))
}
