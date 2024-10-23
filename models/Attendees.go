package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendees struct {
	gorm.Model
	UserID    uuid.UUID `json:"userID" gorm:"type:char(36);index"`
	EventID   uint      `json:"eventID"`
	AppliedAt time.Time `json:"applied_at"`
	Status    string    `json:"status"`
}
