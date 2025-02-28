package contracts

import (
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type AuthenticateUserUseCase interface {
	Execute(input input.AuthenticateUserInput) (string, error)
}
