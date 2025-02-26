package contracts

import (
	"context"

	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, user input.RegisterUserInput) (string, error)
}
