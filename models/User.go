package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string      `json:"id" gorm:"primaryKey;type:char(36)"`
	Username  string      `json:"username" gorm:"index" binding:"required"`
	Name      string      `json:"name"  binding:"required"`
	Email     string      `json:"email" gorm:"unique:index" binding:"required,email"`
	Password  string      `json:"password" binding:"required,min=6"`
	CoverImg  string      `json:"cover_img"`
	Attendees []Attendees `json:"attendees"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
