package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var nextID = 1

func main() {
	r := gin.Default()

	r.POST("/books", createBook)
	r.GET("/books", getBooks)
	// Endpoint untuk mendapatkan buku berdasarkan ID (GET /books/:id)
	r.GET("/books/:id", getBookByID)

	r.Run(":8080")
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// Handler untuk mendapatkan buku berdasarkan ID
func getBookByID(c *gin.Context) {
	// Mendapatkan nilai parameter "id" dari URL
	idStr := c.Param("id")

	// Mengonversi string ID menjadi integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Mencari buku dengan ID yang sesuai
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	// Jika buku tidak ditemukan, kirim respons 404 Not Found
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
