package repositories

import (
	"fmt"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres"
	"github.com/gustavomello-21/pathwaybr-backend/internal/infra/database/postgres/models"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type UserRepository struct {
	client postgres.Client
}

func NewUserRepository(client postgres.Client) repositories.UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (u *UserRepository) FindByEmail(email string) (*entities.User, error) {
	db, err := u.client.Open()
	if err != nil {
		fmt.Println("Error opening database connection: ", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("Error getting database connection: ", err)
		return nil, err
	}
	defer sqlDb.Close()

	var user models.User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		fmt.Println("Bateu aqui")
	}

	entityUser := entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CratedAt:  user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &entityUser, nil
}

func (u *UserRepository) Save(user entities.User) error {
	db, err := u.client.Open()
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

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("Error creating user: ", result.Error)
		return result.Error
	}

	return nil
}
