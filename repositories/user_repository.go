package repositories

import (
	"github.com/google/uuid"
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func GetUser(uuid uuid.UUID) (models.User, error) {
	var user models.User

	if err := db.DB.First(&user, uuid).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}
