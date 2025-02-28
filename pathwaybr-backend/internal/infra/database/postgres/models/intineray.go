package models

import "time"

type Intinerary struct {
	ID        int  `gorm:"primaryKey;autoIncrement;not null"`
	TripID    int  `gorm:"not null"`
	Trip      Trip `gorm:"foreignKey:TripID;references:ID"`
	DayNumber int  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
