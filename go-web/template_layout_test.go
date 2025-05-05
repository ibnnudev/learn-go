package goweb

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	localTemplates, err := template.ParseGlob(filepath.Join("templates", "*.gohtml"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the layout template
	err = localTemplates.ExecuteTemplate(w, "home.gohtml", map[string]any{
		"WebTitle": "Go Web",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	TemplateLayout(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(body))
}
