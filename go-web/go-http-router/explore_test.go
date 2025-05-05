package main

import (
	"fmt"
	"go-http-router/helper"
	"net/http"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

var products = []Product{
	{ID: 1, Name: "Product A", Price: 10.0},
	{ID: 2, Name: "Product B", Price: 20.0},
	{ID: 3, Name: "Product C", Price: 30.0},
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	for index, product := range products {
		fmt.Fprintln(w, "PRODUCT:", index+1)
		fmt.Fprintln(w, "ID:", product.ID)
		fmt.Fprintln(w, "Name:", product.Name)
		fmt.Fprintln(w, "Price:", product.Price)
		fmt.Fprintln(w, "===================================")
	}
}

func getProductByIDHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(w, "DETAIL PRODUCT")
	fmt.Fprintln(w, "Product ID:", id)
	for _, product := range products {
		if fmt.Sprintf("%d", product.ID) == id {
			fmt.Fprintln(w, "ID:", product.ID)
			fmt.Fprintln(w, "Name:", product.Name)
			fmt.Fprintln(w, "Price:", product.Price)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func addProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	newProduct := Product{
		ID:    len(products) + 1,
		Name:  name,
		Price: price,
	}
	products = append(products, newProduct)
	fmt.Fprintln(w, "Product added successfully!")
}

func updateProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	name := r.FormValue("name")
	price, err := helper.StringToFloat(r.FormValue("price"))
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	for index, product := range products {
		if fmt.Sprintf("%d", product.ID) == id {
			products[index].Name = name
			products[index].Price = price
			fmt.Fprintln(w, "Product updated successfully!")
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func setupRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/products", getAllProductsHandler)
	router.GET("/products/:id", getProductByIDHandler)
	router.POST("/product", addProductHandler)
	router.PUT("/product/:id", updateProductHandler)
	return router
}

func TestExplore(t *testing.T) {
	server := http.Server{
		Addr:    ":3000",
		Handler: setupRouter(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
