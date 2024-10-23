package handlers

import (
	"backend/go_backend/models"
	"backend/go_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var events models.Event

		db.Find(&events)
		if err := ctx.ShouldBindJSON(&events); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, events)
	}
}
func GetEvent(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var event models.Event

		if err := db.First(&event, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Event was not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, event)

	}
}
func CreateEvent(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var event models.Event

		err := ctx.ShouldBindJSON(&event)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Parse the date-only fields
		EventDate, err := time.Parse("2006-01-02", event.EventDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event date"})
			return
		}
		// handle file upload
		file, header, err := ctx.Request.FormFile("image")

		if err == nil {
			imageURL, err := utils.SaveImage(file, header)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
				return
			}
			event.CoverImg = imageURL
		}

		// Set applied date for all attendees
		for i := range event.Attendees {
			event.Attendees[i].AppliedAt = time.Now()
			event.Attendees[i].Status = "pending"
		}
		if err := db.Create(&event).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating event"})
			return
		}
		event.EventDate = EventDate.Format("2006-01-02")
		ctx.JSON(http.StatusCreated, event)

	}
}

func DeleteEvent(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var event models.Event

		if err := db.First(&event, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Error occurred!"})
			return

		}

		// Delete associated image
		if err := utils.DeleteImage(event.CoverImg); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image: " + err.Error()})
			return
		}

		if err := db.Delete(&event).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the product"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted Successfully"})

	}
}
func UpdateEvent(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var event models.Event

		if err := db.First(&event, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred!"})
			return

		}
		// bind the json to a new product
		var updatedEvent models.Event

		if err := ctx.ShouldBindJSON(&event); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Handle file upload if a new image is provided
		file, header, err := ctx.Request.FormFile("image")
		if err == nil {
			// New image was provided
			imageURL, err := utils.SaveImage(file, header)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image: " + err.Error()})
				return
			}
			// Delete old image if it exists
			if event.CoverImg != "" {
				if err := utils.DeleteImage(event.CoverImg); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete the image"})
				}
			}
			event.CoverImg = imageURL
		}

		// update the event
		event.Title = updatedEvent.Title
		event.Description = updatedEvent.Description
		event.Occurrence = updatedEvent.Occurrence
		event.Agendas = updatedEvent.Agendas
		event.LocationID = updatedEvent.LocationID
		event.EventDate = updatedEvent.EventDate
		event.Seats = updatedEvent.Seats
		event.Attendees = updatedEvent.Attendees
		event.Type = updatedEvent.Type

		if err := db.Save(&event).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the event"})
		}
		ctx.JSON(http.StatusOK, event)

	}
}
