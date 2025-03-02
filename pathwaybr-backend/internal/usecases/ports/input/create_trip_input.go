package input

import (
	"time"
)

type CreateTripInput struct {
	UserId    int
	StartDate time.Time
	EndDate   time.Time
}
