package entities

import "time"

type User struct {
	// User ID
	ID int

	// User Name
	Username string

	// User Password
	Password string

	// User Email
	Email string

	// Time the user was created
	CratedAt time.Time

	// Last time the user was updated
	UpdatedAt time.Time
}
