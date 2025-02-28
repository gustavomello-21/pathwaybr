package dto

// LoginRequest represents the data needed to authenticate a user.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
