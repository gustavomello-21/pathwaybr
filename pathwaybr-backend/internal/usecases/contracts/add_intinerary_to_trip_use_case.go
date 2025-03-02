package contracts

import "github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"

type AddIntineraryToTripUseCase interface {
	Execute(input input.AddIntineraryToTripInput) error
}
