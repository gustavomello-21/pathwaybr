package entities

import "time"

type Trip struct {
	// Trip ID
	ID int

	// User owner of the trip
	UserID int

	// Time the trip start
	StartDate time.Time

	// Time the trip end
	EndDate time.Time

	// Time the trip was created
	CreatedAt time.Time

	// Last time the trip was updated
	UpdatedAt time.Time
}
