package contracts

import (
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type RegisterUserUseCase interface {
	Execute(user input.RegisterUserInput) (string, error)
}
