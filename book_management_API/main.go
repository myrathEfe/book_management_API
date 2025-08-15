package main

import (
	"book_management_API/initializers"
	"book_management_API/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Book{})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Create endpoint
	r.POST("/books", func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result := initializers.DB.Create(&book)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(201, book)
	})

	//Read endpoint (tüm kitaplar)

	r.GET("/books", func(c *gin.Context) {
		var books []models.Book
		initializers.DB.Find(&books)
		c.JSON(200, books)
	})

	//Read endpoint (tek bir kitap)
	r.GET("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		result := initializers.DB.First(&book, id)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(200, book)
	})

	//Update endpoint

	r.PUT("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		if err := initializers.DB.First(&book, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Book not found"})
			return
		}
		var input models.Book
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		initializers.DB.Model(&book).Updates(input)
		c.JSON(200, book)
	})

	//Delete endpoint

	r.DELETE("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		if err := initializers.DB.First(&book, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Book not found"})
			return
		}
		initializers.DB.Delete(&book)
		c.JSON(200, gin.H{"message": "Book deleted"})
	})

	r.Run()
}
