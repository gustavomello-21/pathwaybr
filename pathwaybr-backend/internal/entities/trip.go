package entities

import "time"

type Trip struct {
	// Trip ID
	ID int `json:"id"`

	// User owner of the trip
	UserID int `json:"user_id"`

	// Time the trip start
	StartDate time.Time `json:"start_date"`

	// Time the trip end
	EndDate time.Time `json:"end_date"`

	// Time the trip was created
	CreatedAt time.Time `json:"created_at"`

	// Last time the trip was updated
	UpdatedAt time.Time `json:"updated_at"`
}
