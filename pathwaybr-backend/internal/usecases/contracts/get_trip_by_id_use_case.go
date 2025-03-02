package contracts

import (
	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type GetTripByIdUseCase interface {
	Execute(input input.GetTripByIdInput) (*entities.Trip, error)
}
