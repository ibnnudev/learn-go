package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ProductItem struct {
	Name  string
	Price float32
}

func (i ProductItem) GetDetail() string {
	return fmt.Sprintf("Item: %s, Price: %.2f", i.Name, i.Price)
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.GetDetail}}`))
	item := ProductItem{Name: "Apple", Price: 1.99}
	err := t.ExecuteTemplate(w, "FUNCTION", item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "/template_function", nil)
	response := httptest.NewRecorder()

	TemplateFunction(response, request)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	item := ProductItem{Name: "Apple", Price: 1.99}
	err := t.ExecuteTemplate(w, "FUNCTION", item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "/template_function", nil)
	response := httptest.NewRecorder()

	TemplateFunctionGlobal(response, request)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateCreateFuncGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	// t = t.Funcs(map[string]any{
	// 	"upper": strings.ToUpper,
	// })

	// function pipeline
	t = t.Funcs(template.FuncMap{
		"upper": strings.ToUpper,
		"sayhello": func(name string) string {
			return fmt.Sprintf("Hello %s", name)
		},
	})

	// t = template.Must(t.Parse(`{{upper .Name}}`))
	// t.ExecuteTemplate(w, "FUNCTION", ProductItem{Name: "Macbook M1", Price: 1.99})

	// function pipeline
	t = template.Must(t.Parse(`{{sayhello .Name | upper}}`))
	err := t.ExecuteTemplate(w, "FUNCTION", ProductItem{Name: "Macbook M1", Price: 1.99})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestTemplateCreateFuncGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "/template_function", nil)
	response := httptest.NewRecorder()

	TemplateCreateFuncGlobal(response, request)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
