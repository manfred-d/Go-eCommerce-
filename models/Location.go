package models

import "gorm.io/gorm"

// Location represents a place where an event takes place

type Location struct {
	gorm.Model
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
