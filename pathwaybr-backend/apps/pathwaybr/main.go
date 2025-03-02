package main

import (
	"fmt"
	"os"

	activity "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/activity/v1"
	auth "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/auth/v1"
	intinerary "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/intinerary/v1"
	trip "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/trip/v1"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/config"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/migrations"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/repositories"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Carregando variáveis de ambiente...")
	err := godotenv.Load("/home/gustavo.melo/Documents/pathwaybr/.env")
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}

	dbConfig := postgres.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
	}

	if err := migrations.Migrate(dbConfig); err != nil {
		fmt.Println("Error migrating models: ", err)
	}

	postgresClient := postgres.NewClient(dbConfig)

	userRepository := repositories.NewUserRepository(*postgresClient)
	tripRepository := repositories.NewTripRepository(*postgresClient)
	intineraryRepository := repositories.NewIntineraryRepository(*postgresClient)
	activityRepository := repositories.NewActivityRepository(*postgresClient)

	authenticateUserUseCase := usecases.NewAuthenticateUserUseCase(userRepository)
	registerUserUseCase := usecases.NewRegisterUserUseCase(userRepository)
	createTripUseCase := usecases.NewCreateTripUseCase(tripRepository)
	getUserTripsUseCase := usecases.NewGetUserTripsUseCase(tripRepository, userRepository)
	getTripByIdUseCae := usecases.NewGetTripByIdUseCase(tripRepository)
	addIntineraryToTripUseCase := usecases.NewAddIntineraryToTripUseCase(intineraryRepository, tripRepository)
	addActivityToIntinerateUseCase := usecases.NewAddActivityToIntineraryUseCase(activityRepository, intineraryRepository)

	controllers := []interface{}{
		auth.NewSessionController(authenticateUserUseCase),
		auth.NewRegisterController(registerUserUseCase),
		trip.NewTripController(createTripUseCase, getUserTripsUseCase, getTripByIdUseCae),
		intinerary.NewIntineraryController(addIntineraryToTripUseCase),
		activity.NewActivityController(addActivityToIntinerateUseCase),
	}

	router := config.Routes(controllers)
	router.Run(":8080")
	fmt.Println("Rodando o servidor...")
}
