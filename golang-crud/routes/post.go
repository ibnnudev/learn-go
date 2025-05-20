package routes

import (
	"golang-crud/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func (r *Router) PostRoutes() {
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}
