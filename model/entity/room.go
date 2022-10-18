package entity

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	ApartmentID string `json:"apartment_id"`
	RoomtypeID  string `json:"roomtype_id"`
	MemberID    string `json:"member_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
