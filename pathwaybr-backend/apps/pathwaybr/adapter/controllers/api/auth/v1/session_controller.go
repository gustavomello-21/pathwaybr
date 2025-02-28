package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/auth/dto"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type SessionController struct {
	authenticationUseCase contracts.AuthenticateUserUseCase
}

func NewSessionController(authenticationUseCase contracts.AuthenticateUserUseCase) *SessionController {
	return &SessionController{
		authenticationUseCase: authenticationUseCase,
	}
}

func (s *SessionController) Create(httpContext *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := httpContext.ShouldBindJSON(&loginRequest); err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := input.AuthenticateUserInput{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
	token, err := s.authenticationUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(200, gin.H{"token": token})

}
