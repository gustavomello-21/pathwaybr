package repositories

import (
	"fmt"
	"time"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/models"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type TripRepository struct {
	client postgres.Client
}

func NewTripRepository(client postgres.Client) repositories.TripRepository {
	return &TripRepository{
		client: client,
	}
}

func (t *TripRepository) FindTripByUserAndDates(userId int, startDate, endDate time.Time) ([]entities.Trip, error) {
	db, err := t.client.Open()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return nil, err
	}
	defer sqlDb.Close()

	var tripsModel []models.Trip

	result := db.Where("user_id = ? AND ((start_date >= ? AND start_date <= ?) OR (end_date >= ? AND end_date <= ?))",
		userId, startDate, startDate, endDate, endDate,
	).Find(&tripsModel)
	if result.Error != nil {
		fmt.Println("Bateu aqui")
	}

	var trips []entities.Trip
	for _, tripModel := range tripsModel {
		tempTripModel := entities.Trip{
			UserID:    tripModel.User.ID,
			StartDate: tripModel.StartDate,
			EndDate:   tripModel.EndDate,
		}
		trips = append(trips, tempTripModel)
	}

	return trips, nil
}

func (t *TripRepository) FindTripsByUserId(userId int) ([]entities.Trip, error) {
	db, err := t.client.Open()
	if err != nil {
		fmt.Println("Failed to open connect: ", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return nil, err
	}
	defer sqlDb.Close()

	var tripsModel []models.Trip
	result := db.Where("user_id = ?", userId).Find(&tripsModel)
	if result.Error != nil {
		fmt.Println("Bateu aqui2")
	}

	var trips []entities.Trip
	for _, tripModel := range tripsModel {
		trips = append(trips, entities.Trip{
			ID:        tripModel.ID,
			UserID:    tripModel.UserID,
			StartDate: tripModel.StartDate,
			EndDate:   tripModel.EndDate,
			CreatedAt: tripModel.CreatedAt,
			UpdatedAt: tripModel.UpdatedAt,
		})
	}

	return trips, nil
}

func (t *TripRepository) FindTripById(tripId int) (*entities.Trip, error) {
	db, err := t.client.Open()
	if err != nil {
		fmt.Println("Failed to open connect: ", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return nil, err
	}
	defer sqlDb.Close()

	var tripModel models.Trip

	result := db.Where("id = ?", tripId).Find(&tripModel)
	if result.Error != nil {
		fmt.Println("Bateu aqui2")
	}

	trip := entities.Trip{
		ID:        tripModel.ID,
		UserID:    tripModel.UserID,
		StartDate: tripModel.StartDate,
		EndDate:   tripModel.EndDate,
		CreatedAt: tripModel.CreatedAt,
		UpdatedAt: tripModel.UpdatedAt,
	}

	return &trip, nil
}

func (t *TripRepository) Save(trip entities.Trip) error {
	db, err := t.client.Open()
	if err != nil {
		fmt.Println("Failed to open connect: ", err)
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return err
	}
	defer sqlDb.Close()

	newTrip := models.Trip{
		UserID:    trip.UserID,
		StartDate: trip.StartDate,
		EndDate:   trip.EndDate,
	}

	result := db.Create(&newTrip)
	if result.Error != nil {
		fmt.Println("Failed to create a trip: ", result.Error)
	}

	return nil
}
