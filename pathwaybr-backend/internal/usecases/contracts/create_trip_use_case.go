package contracts

import "github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"

type CreateTripUseCase interface {
	Execute(input input.CreateTripInput) error
}
