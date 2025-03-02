package repositories

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type ActivityRepository struct {
	client postgres.Client
}

func NewActivityRepository(client postgres.Client) repositories.ActivityRepositories {
	return &ActivityRepository{
		client: client,
	}
}

func (a *ActivityRepository) Save(activity entities.Activity) error {
	db, err := a.client.Open()
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

	newActivity := entities.Activity{
		IntineraryID: activity.IntineraryID,
		Type:         activity.Type,
		Description:  activity.Description,
		StartTime:    activity.StartTime,
		EndTime:      activity.EndTime,
	}
	result := db.Create(&newActivity)
	if result.Error != nil {
		fmt.Println("Failed to create a Activity: ", result.Error)
	}

	return nil
}
