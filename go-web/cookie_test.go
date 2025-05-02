package goweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = "x-useragent"
	cookie.Value = r.URL.Query().Get("useragent")
	cookie.Domain = "localhost"

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Cookie set")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("x-useragent")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Cookie value:", cookie.Value)
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/setcookie", SetCookie)
	mux.HandleFunc("/getcookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}
	defer server.Close()
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "/setcookie?useragent=Mozilla/5.0", nil)
	response := httptest.NewRecorder()
	SetCookie(response, request)

	cookies := response.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "x-useragent" {
			if cookie.Value != "Mozilla/5.0" {
				t.Errorf("Expected cookie value 'Mozilla/5.0', got '%s'", cookie.Value)
			} else {
				t.Logf("Cookie set successfully: %s=%s", cookie.Name, cookie.Value)
			}
			return
		}
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "/getcookie", nil)
	cookie := &http.Cookie{
		Name:  "x-useragent",
		Value: "Mozilla/5.0",
	}
	request.AddCookie(cookie)

	response := httptest.NewRecorder()
	GetCookie(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	expected := "Cookie value:Mozilla/5.0"
	if response.Body.String() != expected {
		t.Errorf("Expected body '%s', got '%s'", expected, response.Body.String())
	}
}
