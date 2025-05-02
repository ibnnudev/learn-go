package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PageTemplate struct {
	Title      string
	Name       string
	Address    map[string]any
	IsActive   bool
	FinalGrade float32
}

func TemplateAction(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	templ.ExecuteTemplate(w, "simple.gohtml", PageTemplate{
		Title: "Hello, World!",
		Name:  "John Doe",
		Address: map[string]any{
			"Street":  "123 Main St",
			"City":    "Anytown",
			"State":   "CA",
			"ZipCode": "12345",
		},
		IsActive:   true,
		FinalGrade: 3.0,
	})
}

func TestTemplateAction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateAction(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
