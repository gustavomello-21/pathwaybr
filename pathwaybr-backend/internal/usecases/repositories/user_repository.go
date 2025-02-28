package repositories

import (
	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	Save(user entities.User) error
}
