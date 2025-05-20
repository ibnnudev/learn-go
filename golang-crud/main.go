package main

import (
	"golang-crud/database"
	"golang-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDB()

	postRoutes := routes.Router{Engine: router}

	postRoutes.PostRoutes()

	router.Run(":8080")
}
