package controllers

import (
	"kasir-backend/database"
	"kasir-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create transaction
func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate total
	var product models.Product
	database.DB.First(&product, input.ProductId)
	input.Total = float64(input.Quantity) * product.Price

	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// Get transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
