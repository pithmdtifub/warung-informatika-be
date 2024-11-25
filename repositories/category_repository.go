package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/helpers"
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
	res := db.DB.Save(category)
	if err := helpers.CheckRowsAffected(res.RowsAffected); err != nil {
		return err
	}

	return res.Error
}

func DeleteCategory(id uint) error {
	category := models.Category{ID: id}
	res := db.DB.Delete(&category)

	if err := helpers.CheckRowsAffected(res.RowsAffected); err != nil {
		return err
	}

	return res.Error
}
