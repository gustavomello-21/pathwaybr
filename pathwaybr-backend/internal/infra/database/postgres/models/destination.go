package models

import "time"

type Destination struct {
	ID        int    `gorm:"primaryKey;autoIncrement;not null"`
	TripID    int    `gorm:"not null"`
	Trip      Trip   `gorm:"foreignKey:TripID;references:ID"`
	City      string `gorm:"not null"`
	Country   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
