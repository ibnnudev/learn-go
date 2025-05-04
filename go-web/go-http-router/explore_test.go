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

func TestExplore(t *testing.T) {
	var products []Product = []Product{
		{ID: 1, Name: "Product A", Price: 10.0},
		{ID: 2, Name: "Product B", Price: 20.0},
		{ID: 3, Name: "Product C", Price: 30.0},
	}

	router := httprouter.New()
	router.GET("/products", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for index, product := range products {
			fmt.Fprintln(w, "PRODUCT:", index+1)
			fmt.Fprintln(w, "ID:", product.ID)
			fmt.Fprintln(w, "Name:", product.Name)
			fmt.Fprintln(w, "Price:", product.Price)
			fmt.Fprintln(w, "===================================")
		}
	})
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		fmt.Fprintln(w, "DETAIL PRODUCT")
		fmt.Fprintln(w, "Product ID:", id)
		for _, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				fmt.Fprintln(w, "ID:", product.ID)
				fmt.Fprintln(w, "Name:", product.Name)
				fmt.Fprintln(w, "Price:", product.Price)
				break
			}
		}
	})
	router.POST("/product", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	})

	router.PUT("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
