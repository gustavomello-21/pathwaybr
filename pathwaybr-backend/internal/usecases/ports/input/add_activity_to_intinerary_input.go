package input

import "time"

type AddActivityToIntineraryInput struct {
	IntinerarieId int
	Type          string
	Description   string
	StartTime     time.Time
	EndTime       time.Time
}
