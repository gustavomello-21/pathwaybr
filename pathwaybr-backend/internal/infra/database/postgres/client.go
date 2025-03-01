package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
}

type Client struct {
	dataSourceName string
}

func NewClient(config DatabaseConfig) *Client {
	fmt.Println(config.Password)
	return &Client{
		dataSourceName: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			config.Host,
			config.User,
			config.Password,
			config.Dbname,
			config.Port,
			config.Sslmode,
		),
	}
}

func (c *Client) Open() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Println("Bateu aqui")
		fmt.Println("error opening database connection: ", err)
		return nil, err
	}

	return db, nil
}
