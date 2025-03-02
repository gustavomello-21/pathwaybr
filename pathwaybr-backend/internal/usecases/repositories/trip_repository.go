package repositories

import (
	"time"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
)

type TripRepository interface {
	FindTripByUserAndDates(userId int, startDate, endDate time.Time) ([]entities.Trip, error)
	FindTripsByUserId(userId int) ([]entities.Trip, error)
	FindTripById(tripId int) (*entities.Trip, error)
	Save(trip entities.Trip) error
}
