package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string      `json:"title" gorm:"index" binding:"required"`
	Occurrence  string      `json:"is_Series" binding:"required"`
	Description string      `json:"description"`
	CoverImg    string      `json:"cover_img" binding:"required"`
	Agendas     string      `json:"agendas" binding:"required"`
	LocationID  uint        `json:"location_id"`
	Location    Location    `json:"location"`
	EventDate   string      `json:"event_time"`
	Seats       int         `json:"no_of_attendees"`
	Attendees   []Attendees `json:"attendees"`
	Type        int         `json:"event_type" binding:"required"`
	HostedByID  uuid.UUID   `json:"hostedby_id" gorm:"type:char(36);index"`
}
