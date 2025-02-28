package models

import "time"

type Activity struct {
	ID           uint       `gorm:"primaryKey;autoIncrement;not null"`
	IntineraryID uint       `gorm:"not null"`
	Intinerary   Intinerary `gorm:"foreignKey:IntineraryID;references:ID"`
	Type         string     `gorm:"not null"`
	Description  string     `gorm:"not null"`
	StartTime    string     `gorm:"not null"`
	EndTime      string     `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
