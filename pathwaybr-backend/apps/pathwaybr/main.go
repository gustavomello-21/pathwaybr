package main

import (
	"fmt"
	"os"

	v1 "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/auth/v1"
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
	authenticateUserUseCase := usecases.NewAuthenticateUserUseCase(userRepository)
	controllers := []interface{}{
		v1.NewSessionController(authenticateUserUseCase),
	}

	router := config.Routes(controllers)
	router.Run(":8080")
	fmt.Println("Rodando o servidor...")
}
