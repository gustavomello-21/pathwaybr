package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/auth/dto"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type RegisterController struct {
	RegisterUserUseCase contracts.RegisterUserUseCase
}

func NewRegisterController(registerUserUseCase contracts.RegisterUserUseCase) *RegisterController {
	return &RegisterController{
		RegisterUserUseCase: registerUserUseCase,
	}
}

func (r *RegisterController) Create(httpContext *gin.Context) {
	var registerRequestDto dto.RegisterRequestDto
	if err := httpContext.ShouldBindJSON(&registerRequestDto); err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := input.RegisterUserInput{
		Username: registerRequestDto.Username,
		Email:    registerRequestDto.Email,
		Password: registerRequestDto.Password,
	}

	token, err := r.RegisterUserUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(201, gin.H{"Token": token})
}
