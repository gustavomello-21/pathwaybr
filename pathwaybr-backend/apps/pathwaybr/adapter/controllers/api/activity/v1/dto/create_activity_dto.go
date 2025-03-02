package dto

type CreateActivityDto struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
