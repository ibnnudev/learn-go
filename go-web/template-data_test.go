package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Page struct {
	Title   string
	Name    string
	Address map[string]any
}

func TemplateData(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", Page{
		Title: "Hello, World!",
		Name:  "John Doe",
		Address: map[string]any{
			"Street":  "123 Main St",
			"City":    "Anytown",
			"State":   "CA",
			"ZipCode": "12345",
		},
	})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateData(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
