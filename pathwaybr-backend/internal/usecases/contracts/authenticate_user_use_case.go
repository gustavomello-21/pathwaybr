package contracts

import (
	"context"

	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type AuthenticateUserUseCase interface {
	Execute(ctx context.Context, input input.AuthenticateUserInput) (string, error)
}
