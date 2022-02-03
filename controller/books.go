package controller

import (
	"net/http"
	"simpleLoadbalancer/model"

	"github.com/gin-gonic/gin"
)

// It should have pagination and it is not very scalble now
func FindBooks(c *gin.Context) {
	var books []model.Book
	model.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	var book model.Book
	// we use connection pool (openConnection:100, maxIdelConnection:10, connctionLife:1h)
	if err := model.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
