package input

type AuthenticateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
