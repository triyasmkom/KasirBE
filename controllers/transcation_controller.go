package controllers

import (
	"kasir-backend/database"
	"kasir-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create transaction
func CreateTransaction(c *gin.Context) {

	var input struct {
		Cart []struct {
			Produk string  `json:"produk"`
			Harga  float64 `json:"harga"`
		}
		TotalHarga       float64 `json:"total_harga"`
		MetodePembayaran string  `json:"metode_pembayaran"`
		WaktuTransaksi   string  `json:"waktu_transaksi"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Konversi userId dari float64 ke uint
	userIdUint := uint(userId.(float64))

	waktuTransaksi, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", input.WaktuTransaksi+"+07:00")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction time format"})
		return
	}

	transaction := models.Transaction{
		UserId:           userIdUint,
		TotalHarga:       input.TotalHarga,
		MetodePembayaran: input.MetodePembayaran,
		WaktuTransaksi:   waktuTransaksi,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	for _, item := range input.Cart {
		transactionItem := models.TransactionItem{
			TransactionId: transaction.ID,
			ProductName:   item.Produk,
			Price:         item.Harga,
		}

		if err := database.DB.Create(&transactionItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction item"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
		"data":    input,
	})
}

// Get transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
