package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
	"github.com/gustavomello-21/pathwaybr-backend/internal/utils/hasher"
	"github.com/gustavomello-21/pathwaybr-backend/internal/utils/jwt"
)

type RegisterUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewRegisterUserUseCase(userRepository repositories.UserRepository) contracts.RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, input input.RegisterUserInput) (string, error) {
	user, err := uc.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		fmt.Println("error finding user by email: ", err)
		return "", err
	}

	if user.ID != 0 {
		return "", errors.New("user already exists")
	}

	// TODO: Implement This
	hashedPassword, err := hasher.HashPassword(input.Password)
	if err != nil {
		fmt.Println("error hashing password: ", err)
		return "", err
	}

	user = entities.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: hashedPassword,
	}

	err = uc.userRepository.Save(ctx, user)
	if err != nil {
		fmt.Println("error saving user: ", err)
		return "", err
	}

	// TODO: implement this
	token := jwt.GenerateToken()

	return token, nil
}
