package handlers

import (
	"backend/go_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAttendee(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var attendee models.Attendees

		if err := ctx.ShouldBindJSON(&attendee); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// set the application time and status
		attendee.AppliedAt = time.Now()
		attendee.Status = "Pending"

		// insert attendee data
		if err := db.Create(&attendee).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up for the event"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Signed successfully", "attendee": attendee})
	}
}

// remove attending details
func RemoveAttendee(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var attendee models.Attendees

		//extract user id and event id from query params
		userID := ctx.Query("userID")
		eventID := ctx.Query("eventID")

		//find the attendee record
		if err := db.Where("userID  = ? AND eventID = ?", userID, eventID).First(&attendee).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Attendee Not Found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// delete records
		if err := db.Delete(&attendee).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel the event"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Attendance cancelled successfully"})

	}
}
