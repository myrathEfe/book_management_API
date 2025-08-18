package routes

import (
	"book_management_API/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// CRUD routes
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
