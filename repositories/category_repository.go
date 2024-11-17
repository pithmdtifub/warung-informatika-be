package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Find(&categories).Error

	return categories, err
}

func GetCategory(id int) (models.Category, error) {
	category := models.Category{ID: uint(id)}
	err := db.DB.Find(&category).Error

	return category, err
}

func CreateCategory(category *models.Category) error {
	return db.DB.Create(category).Error
}
