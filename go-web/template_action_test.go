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

type Item struct {
	Name  string
	Price float32
}

type PriceListTemplate struct {
	ID         int
	Items      []Item
	TotalItems int
	TotalPrice float32
}

func PriceListAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/pricelist.gohtml"))
	items := []Item{
		{"Lenovo ThinkPad", 1200.00},
		{"Dell XPS", 1500.00},
		{"MacBook Pro", 2000.00},
		{"Asus ZenBook", 1300.00},
		{"HP Spectre", 1400.00},
	}

	var totalPrice float32
	for _, item := range items {
		totalPrice += item.Price
	}

	tpl.ExecuteTemplate(w, "pricelist.gohtml", PriceListTemplate{
		ID:         1,
		Items:      items,
		TotalItems: len(items),
		TotalPrice: totalPrice,
	})

}

func TestPriceListAction(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/pricelist", PriceListAction)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	defer server.Close()
}
