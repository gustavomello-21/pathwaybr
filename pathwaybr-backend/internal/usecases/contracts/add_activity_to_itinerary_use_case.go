package contracts

import "github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"

type AddActivityToIntineraryUseCase interface {
	Execute(input input.AddActivityToIntineraryInput) error
}
