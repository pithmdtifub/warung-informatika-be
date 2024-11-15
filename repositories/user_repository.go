package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

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
