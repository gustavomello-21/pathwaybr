package models

import "time"

type Activity struct {
	ID           int        `gorm:"primaryKey;autoIncrement;not null"`
	IntineraryID int        `gorm:"not null"`
	Intinerary   Intinerary `gorm:"foreignKey:IntineraryID;references:ID"`
	Type         string     `gorm:"not null"`
	Description  string     `gorm:"not null"`
	StartTime    time.Time  `gorm:"not null"`
	EndTime      time.Time  `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
