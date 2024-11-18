package seeders

import (
	"github.com/google/uuid"
	"log"
	"warung-informatika-be/helpers"
	"warung-informatika-be/models"
	repo "warung-informatika-be/repositories"
)

func UserSeeder(user models.User) {
	password, _ := helpers.HashPassword(user.Password)

	user.ID, _ = uuid.NewV7()
	user.Password = password

	if err := repo.CreateUser(&user); err != nil {
		log.Fatal("Failed to create user")
	}
}
