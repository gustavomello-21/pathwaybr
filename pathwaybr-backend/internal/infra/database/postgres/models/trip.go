package models

import "time"

type Trip struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
