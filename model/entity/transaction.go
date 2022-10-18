package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Status    bool           `json:"status"`
	Evalue    float64        `json:"e_value" gorm:"type:decimal(10,2)"`
	Wvalue    float64        `json:"w_value" gorm:"type:decimal(10,2)"`
	RoomID    uint           `json:"room_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
