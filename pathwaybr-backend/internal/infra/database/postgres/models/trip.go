package models

import "time"

type Trip struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	StartDate string `gorm:"not null"`
	EndDate   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
