package contracts

import (
	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type GetUserTripsUseCase interface {
	Execute(input input.GetUserTripsInput) ([]entities.Trip, error)
}
