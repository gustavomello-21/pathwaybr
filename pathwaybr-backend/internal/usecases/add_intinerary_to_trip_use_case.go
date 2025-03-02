package usecases

import (
	"errors"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type AddIntineraryToTripUseCase struct {
	intineraryRepository repositories.IntineraryRepository
	tripRepository       repositories.TripRepository
}

func NewAddIntineraryToTripUseCase(
	intineraryRepository repositories.IntineraryRepository,
	tripRepository repositories.TripRepository,
) contracts.AddIntineraryToTripUseCase {
	return &AddIntineraryToTripUseCase{
		intineraryRepository: intineraryRepository,
		tripRepository:       tripRepository,
	}
}

func (uc *AddIntineraryToTripUseCase) Execute(input input.AddIntineraryToTripInput) error {
	trip, err := uc.tripRepository.FindTripById(input.TripId)
	if err != nil {
		return err
	}

	if trip == nil {
		return errors.New("Trip Not Found")
	}

	intinerary := entities.Intinerary{
		TripID:    input.TripId,
		DayNumber: input.DayNumber,
	}

	if err := uc.intineraryRepository.Save(intinerary); err != nil {
		return err
	}

	return nil
}
