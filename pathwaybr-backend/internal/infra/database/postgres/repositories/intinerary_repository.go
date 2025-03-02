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
