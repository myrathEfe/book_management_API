package main

import (
	"book_management_API/initializers"
	"book_management_API/models"
	"book_management_API/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()

	err := initializers.DB.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Failed to migrate database schema: " + err.Error())
	}
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
