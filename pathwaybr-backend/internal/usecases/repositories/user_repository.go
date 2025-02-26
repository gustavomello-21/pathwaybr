package repositories

import (
	"context"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (entities.User, error)
	Save(ctx context.Context, user entities.User) error
}
