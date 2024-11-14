package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetUser(id any, username any) (models.User, error) {
	var user models.User
	var err error

	if id != nil {
		err = db.DB.First(&user, id).Error
	} else if username != nil {
		err = db.DB.Where("username = ?", username).First(&user).Error
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
