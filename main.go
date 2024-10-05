package main

import (
	"kasir-backend/database"
	"kasir-backend/routers"
)

func main() {
	// Koneksi ke database
	database.ConnectDatabase()

	// Setup router
	r := routers.SetupRouter()

	// Jalankan server
	r.Run(":8080")
}
