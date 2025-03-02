package dto

type CreateIntineraryDto struct {
	TripId    int `json:"trip_id"`
	DayNumber int `json:"day_number"`
}
