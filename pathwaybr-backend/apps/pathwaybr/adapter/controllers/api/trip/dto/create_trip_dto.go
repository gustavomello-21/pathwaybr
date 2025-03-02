package dto

type CreateTripDto struct {
	UserId    int    `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
