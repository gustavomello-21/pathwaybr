package entities

import "time"

type Destination struct {
	// Destination ID
	ID int

	// Id of the trip that the destination belongs to
	TripID int

	// City of the destination
	City string

	// Country of the destination
	Country string

	// Time the destination was created
	CreatedAt time.Time

	// Last time the destination was updated
	UpdatedAt time.Time
}
