package entity

import (
	"time"

	"gorm.io/gorm"
)

type Apartment struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name"`
	Branch   string         `json:"branch"`
	Phone    string         `json:"phoen"`
	Adress   string         `json:"adress"`
	UserID   uint           `json:"user_id"`
	Image    string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
