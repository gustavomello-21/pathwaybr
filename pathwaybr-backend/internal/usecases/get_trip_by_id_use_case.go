package usecases

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type GetTripByIdUseCase struct {
	tripRepository repositories.TripRepository
}

func NewGetTripByIdUseCase(tripRepository repositories.TripRepository) contracts.GetTripByIdUseCase {
	return &GetTripByIdUseCase{
		tripRepository: tripRepository,
	}
}

func (uc *GetTripByIdUseCase) Execute(input input.GetTripByIdInput) (*entities.Trip, error) {
	trip, err := uc.tripRepository.FindTripById(input.TripId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return trip, nil
}
