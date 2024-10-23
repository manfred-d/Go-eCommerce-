package handlers

import (
	"backend/go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get a single location
func GetLocation(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var location []models.Location

		if err := db.First(&location, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}
		ctx.JSON(http.StatusOK, location)
	}
}

// get all locations
func Getlocations(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var locations []models.Location

		db.Find(&locations)
		ctx.JSON(http.StatusOK, locations)
	}
}

func Createlocation(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var location models.Location

		err := ctx.ShouldBindJSON(&location)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		if err := db.Create(&location).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, location)
	}
}

// delete a location
func Deletelocation(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var location models.Location

		if err := db.First(&location, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}

		if err := db.Delete(&location).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the location"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "location deleted Successfully"})
	}
}

// update a location
func Updatelocation(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var location models.Location

		if err := db.First(&location, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}

		// bind the json to a new location
		var updatedProd models.Location

		if err := ctx.ShouldBindJSON(&updatedProd); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update the location
		location.Address = updatedProd.Address
		location.City = updatedProd.City
		location.State = updatedProd.State
		location.Country = updatedProd.Country

		if err := db.Save(&location).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the location"})
		}
		ctx.JSON(http.StatusOK, location)
	}
}
