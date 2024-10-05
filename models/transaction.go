package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ProductId uint    `json:"product_id"`
	UserId    uint    `json:"user_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
