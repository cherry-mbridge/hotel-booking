package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	UUIDModel
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	RoomID    uuid.UUID `gorm:"type:char(36)" json:"room_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Timestamps
}
