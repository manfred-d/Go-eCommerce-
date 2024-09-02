package routes

import (
	"backend/go_backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	users := router.Group("/api/users")
	{
		users.GET("/", handlers.GetUsers(db))
		users.POST("/", handlers.CreateUser(db))
		users.GET("/:id", handlers.GetUser(db))
		users.POST("/login", handlers.SignIn(db))
		users.PUT("/:id", handlers.UpdateUser(db))
		users.DELETE("/:id", handlers.DeleteUser(db))
	}

	products := router.Group("/api/products")
	{
		products.GET("/", handlers.GetProducts(db))
		products.POST("/", handlers.CreateProduct(db))
		products.GET("/:id", handlers.GetProduct(db))
		products.DELETE("/:id", handlers.DeleteProduct(db))
		products.PUT("/:id", handlers.UpdateProduct(db))
	}
}
