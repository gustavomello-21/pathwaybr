package main

import (
	"fmt"
	"os"

	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/config"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/migrations"
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

	router := config.Routes()
	router.Run(":8080")
	fmt.Println("Rodando o servidor...")
}
