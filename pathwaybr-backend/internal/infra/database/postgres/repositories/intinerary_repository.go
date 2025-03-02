package repositories

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/models"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type IntineraryRepository struct {
	client postgres.Client
}

func NewIntineraryRepository(client postgres.Client) repositories.IntineraryRepository {
	return &IntineraryRepository{
		client: client,
	}
}

func (i *IntineraryRepository) FindById(intineraryId int) (*entities.Intinerary, error) {
	db, err := i.client.Open()
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

	var intineraryModel models.Intinerary
	result := db.Where("id = ?", intineraryId).Find(&intineraryModel)
	if result.Error != nil {
		return nil, err
	}

	intinerary := entities.Intinerary{
		ID:        intineraryModel.ID,
		TripID:    intineraryModel.TripID,
		DayNumber: intineraryModel.DayNumber,
		CreatedAt: intineraryModel.CreatedAt,
		UpdatedAt: intineraryModel.UpdatedAt,
	}

	return &intinerary, nil
}

func (i *IntineraryRepository) Save(intinerary entities.Intinerary) error {
	db, err := i.client.Open()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Erro while oppenin database connection: ", err)
		return err
	}
	defer sqlDb.Close()

	newIntinerary := models.Intinerary{
		TripID:    intinerary.TripID,
		DayNumber: intinerary.DayNumber,
	}

	result := db.Create(&newIntinerary)
	if result.Error != nil {
		fmt.Println("Failed to create a trip: ", result.Error)
	}

	return nil
}
