package entities

import "time"

type Intinerary struct {
	// Intinerary ID
	ID int

	// Id of the trip that the intinerary belongs to
	TripID int

	// Day number of the trip
	DayNumber int

	// Time the intinerary was created
	CreatedAt time.Time

	// Last time the intinerary was updated
	UpdatedAt time.Time
}
