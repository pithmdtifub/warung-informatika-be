package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Order("id").Find(&categories).Error

	return categories, err
}

func GetCategory(id uint) (models.Category, error) {
	var category models.Category
	err := db.DB.First(&category, id).Error

	return category, err
}

func CreateCategory(category *models.Category) error {
	return db.DB.Create(category).Error
}

func UpdateCategory(category *models.Category) error {
	return db.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	category := models.Category{ID: id}
	err := db.DB.Delete(&category).Error

	return err
}
