package entities

import "time"

type Activity struct {
	// Activity ID
	ID int

	// Id of the intinerary that the activity belongs to
	IntineraryID int

	// Activity Type (Ex: "voo", "hotel", "passeio")
	Type string

	// Description of the activity
	Description string

	// Time the activity starts
	StartTime time.Time

	// Time the activity ends
	EndTime time.Time

	// Time the activity was created
	CreatedAt time.Time

	// Last time the activity was updated
	UpdatedAt time.Time
}
