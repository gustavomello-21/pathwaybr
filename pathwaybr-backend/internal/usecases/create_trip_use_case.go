package usecases

import (
	"errors"
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type CreateTripUseCase struct {
	tripRepository repositories.TripRepository
}

func NewCreateTripUseCase(tripRepository repositories.TripRepository) contracts.CreateTripUseCase {
	return &CreateTripUseCase{
		tripRepository: tripRepository,
	}
}

func (uc *CreateTripUseCase) Execute(input input.CreateTripInput) error {
	trips, err := uc.tripRepository.FindTripByUserAndDates(input.UserId, input.StartDate, input.EndDate)
	if err != nil {
		fmt.Println("Erro to get trips by user and dates: ", err)
		return err
	}

	if len(trips) != 0 {
		return errors.New("already has a trip to this period")
	}

	trip := entities.Trip{
		UserID:    input.UserId,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}

	if err := uc.tripRepository.Save(trip); err != nil {
		fmt.Println("Failed to save a trip")
		return err
	}

	return nil
}
