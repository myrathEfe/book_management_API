package controllers

import (
	"book_management_API/initializers"
	"book_management_API/models"
	"github.com/gin-gonic/gin"
)

// Create endpoint
func CreateBook(c *gin.Context) {
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
}

//Read endpoint (tüm kitaplar)

func GetBooks(c *gin.Context) {
	var books []models.Book
	initializers.DB.Find(&books)
	c.IndentedJSON(200, books)
}

// Read endpoint (tek bir kitap)
func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	result := initializers.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, book)
}

//Update endpoint

func UpdateBook(c *gin.Context) {
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
}

//Delete endpoint

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := initializers.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	initializers.DB.Delete(&book)
	c.JSON(200, gin.H{"message": "Book deleted"})
}
