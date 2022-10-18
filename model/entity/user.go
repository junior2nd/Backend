package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Email    string         `json:"email"`
	Password string         `json:"-" gorm:"comlumn:password"`
	Name     string         `json:"name"`
	Adress   string         `json:"adress"`
	Phone    string         `json:"phone"`
	Image    string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
