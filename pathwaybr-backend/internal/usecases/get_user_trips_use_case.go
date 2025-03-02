package usecases

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type GetUserTripsUseCase struct {
	tripRepository repositories.TripRepository
	userRepository repositories.UserRepository
}

func NewGetUserTripsUseCase(
	tripRepository repositories.TripRepository,
	userRepositories repositories.UserRepository,
) contracts.GetUserTripsUseCase {
	return &GetUserTripsUseCase{
		tripRepository: tripRepository,
		userRepository: userRepositories,
	}
}

func (uc *GetUserTripsUseCase) Execute(input input.GetUserTripsInput) ([]entities.Trip, error) {
	trips, err := uc.tripRepository.FindTripsByUserId(input.UserId)
	if err != nil {
		fmt.Println("erro while getting trips by user id: ", err)
		return nil, err
	}

	return trips, nil
}
