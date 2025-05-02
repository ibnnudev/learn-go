package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed templates/*
var localTemplates embed.FS
var templateEmbed = template.Must(template.ParseFS(localTemplates, "templates/*"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	templateEmbed.ExecuteTemplate(w, "simple.gohtml", nil)
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	TemplateCaching(response, request)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "text/html; charset=utf-8", response.Header().Get("Content-Type"))

}
