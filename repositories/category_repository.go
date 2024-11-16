package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetCategories(limit, offset int) ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Limit(limit).Offset(offset).Find(&categories).Error

	return categories, err
}

func GetCategory(id int) error {
	var category models.Category
	err := db.DB.Find(&category, id).Error

	return err
}

func CreateCategory(category *models.Category) error {
	return db.DB.Create(category).Error
}
