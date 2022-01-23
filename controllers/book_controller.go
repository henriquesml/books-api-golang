package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/henriquesml/go-api/database"
	"github.com/henriquesml/go-api/models"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func ShowBooks(c *gin.Context) {

	db := database.GetDatabase()

	var books models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}
