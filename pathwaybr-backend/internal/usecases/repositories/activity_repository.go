package repositories

import "github.com/gustavomello-21/pathwaybr-backend/internal/entities"

type ActivityRepositories interface {
	Save(entities.Activity) error
}
