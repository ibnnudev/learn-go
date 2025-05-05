package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	templateEmbed.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Template Auto Escaping",
		// "Body":  "<script>alert('XSS')</script>",
		// if u want to disable auto escaping, you can use template.HTML or based on the script that u want to run
		// "Body": template.HTML("<h1>Template Auto Escaping</h1><script>alert('XSS')</script>"),
		"Body": template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	TemplateAutoEscape(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started on :8080")
	}
}

func TestTemplateXSS(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started on :8080")
	}
}
