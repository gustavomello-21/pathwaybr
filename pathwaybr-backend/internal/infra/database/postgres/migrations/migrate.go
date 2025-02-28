package migrations

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/models"
)

func Migrate(dbConfig postgres.DatabaseConfig) error {
	db, err := postgres.NewClient(dbConfig).Open()
	if err != nil {
		fmt.Println("Error opening database connection: ", err)
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Error getting database connection: ", err)
		return err
	}
	defer sqlDb.Close()

	err = db.AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.Intinerary{},
		&models.Destination{},
		&models.Activity{},
	)
	if err != nil {
		fmt.Println("Error migrating models: ", err)
		return err
	}

	return nil
}
