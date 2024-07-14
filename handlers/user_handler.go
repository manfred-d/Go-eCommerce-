package handlers

import (
	"backend/go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users models.User

		db.Find(&users)
		ctx.JSON(http.StatusOK, users)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}

// get a singe user
func GetUser(db *gorm.DB) gin.HandlerFunc {
	// Implementation
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user []models.User

		if err := db.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred"})
			return
		}
		ctx.JSON(http.StatusOK, user)

	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user models.User

		if err := db.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return
		}

		var updatedUser models.User
		if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Username = updatedUser.Username
		user.Email = updatedUser.Email
		user.Password = updatedUser.Password

		if err := db.Save(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the user"})
		}
		ctx.JSON(http.StatusOK, user)

	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user models.User

		if err := db.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"Error": "User Not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}
		ctx.JSON(http.StatusOK, gin.H{"message": "user successfully deleted"})

	}
}
