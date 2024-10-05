package controllers

import (
	"kasir-backend/database"
	"kasir-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
