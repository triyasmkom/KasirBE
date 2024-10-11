package routers

import (
	"kasir-backend/auth"
	"kasir-backend/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ubah sesuai dengan origin frontend kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// User routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Use(auth.AuthMiddleware())
	// User routes (protected)
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.AddUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id", controllers.EditUser)

	// Product routes (protected)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.PUT("/products/:id", controllers.EditProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Transaction routes (protected)
	r.POST("/transactions", controllers.CreateTransaction)
	r.GET("/transactions", controllers.GetTransactions)

	return r
}
