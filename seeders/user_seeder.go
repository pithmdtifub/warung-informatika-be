package seeders

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"warung-informatika-be/helpers"
	"warung-informatika-be/models"
	repo "warung-informatika-be/repositories"
)

func UserSeeder(user models.User) {
	_, err := repo.GetUserByUsername(user.Username)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal("Username already exist")
	}

	password, _ := helpers.HashPassword(user.Password)

	user.ID, _ = uuid.NewV7()
	user.Password = password

	if err := repo.CreateUser(&user); err != nil {
		log.Fatal("Failed to create user")
	}
}
