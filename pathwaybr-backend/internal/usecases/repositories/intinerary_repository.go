package repositories

import "github.com/gustavomello-21/pathwaybr-backend/internal/entities"

type IntineraryRepository interface {
	FindById(intinenaryId int) (*entities.Intinerary, error)
	Save(intinerary entities.Intinerary) error
}
