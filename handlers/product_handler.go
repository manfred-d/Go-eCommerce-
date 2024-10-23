package handlers

import (
	"backend/go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get a single product
func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product []models.Product

		if err := db.First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}
		ctx.JSON(http.StatusOK, product)
	}
}

// get all products
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product

		db.Find(&products)
		ctx.JSON(http.StatusOK, products)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		if err := db.Create(&product).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
			return
		}
		ctx.JSON(http.StatusCreated, product)
	}
}

// delete a product
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}

		if err := db.Delete(&product).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the product"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted Successfully"})
	}
}

// update a product
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}

		// bind the json to a new product
		var updatedProd models.Product

		if err := ctx.ShouldBindJSON(&updatedProd); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update the product
		product.Name = updatedProd.Name
		product.Description = updatedProd.Description
		product.Price = updatedProd.Price

		if err := db.Save(&product).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the product"})
		}
		ctx.JSON(http.StatusOK, product)
	}
}
