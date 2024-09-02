package handlers

import (
	"backend/go_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users models.User

		db.Find(&users)
		ctx.JSON(http.StatusOK, users)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		// Generate a new UUID
		user.ID = uuid.NewString()

		// Generate JWT token
		token, err := generateJWTToken(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Hash the password
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = hashedPassword

		if err := db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User created successful",
			"token":   token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		})
	}
}

// sign up - login
func SignIn(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var LoginCreds struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}
		if err := ctx.ShouldBindJSON(&LoginCreds); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var user models.User
		if err := db.Where("Email = ?", &LoginCreds.Email).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user"})
			}
			return
		}
		// Verify password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginCreds.Password)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		// Generate JWT token
		token, err := generateJWTToken(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Login Successful",
			"token":   token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		})

	}
}

// get a singe user
func GetUser(db *gorm.DB) gin.HandlerFunc {
	// Implementation
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user []models.User

		if err := db.First(&user, &id).Error; err != nil {
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

func generateJWTToken(userID string) (string, error) {
	// Set up your JWT secret key
	secretKey := []byte("1232")

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	return token.SignedString(secretKey)
}
