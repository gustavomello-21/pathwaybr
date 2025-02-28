package usecases

import (
	"errors"

	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
	"github.com/gustavomello-21/pathwaybr-backend/internal/utils/hasher"
	"github.com/gustavomello-21/pathwaybr-backend/internal/utils/jwt"
)

type AuthenticateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewAuthenticateUserUseCase(userRepository repositories.UserRepository) contracts.AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *AuthenticateUserUseCase) Execute(input input.AuthenticateUserInput) (string, error) {
	user, err := uc.userRepository.FindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("user not found")
	}

	valid := hasher.CompareHashAndPassword(user.Password, input.Password)

	if !valid {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}
