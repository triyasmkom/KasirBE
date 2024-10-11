package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId           uint              `json:"user_id"`
	TotalHarga       float64           `json:"total_harga"`
	MetodePembayaran string            `json:"metode_pembayaran"`
	WaktuTransaksi   time.Time         `json:"waktu_transaksi"`
	Items            []TransactionItem `json:"items" gorm:"foreignKey:TransactionId"`
}

type TransactionItem struct {
	gorm.Model
	TransactionId uint    `json:"transaction_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
}
