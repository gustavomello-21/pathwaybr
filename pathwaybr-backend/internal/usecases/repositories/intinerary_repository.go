package repositories

import "github.com/gustavomello-21/pathwaybr-backend/internal/entities"

type IntineraryRepository interface {
	Save(intinerary entities.Intinerary) error
}
