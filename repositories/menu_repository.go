package repositories

import (
	"gorm.io/gorm/clause"
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
)

func GetMenus() ([]models.Menu, error) {
	var menus []models.Menu
	err := db.DB.Preload(clause.Associations).Find(&menus).Error

	return menus, err
}

func GetMenu(id int) (models.Menu, error) {
	var menu models.Menu
	err := db.DB.Preload(clause.Associations).First(&menu, id).Error

	return menu, err
}

func GetMenusByCategoryId(id int) ([]models.Menu, error) {
	var menus []models.Menu
	err := db.DB.Preload(clause.Associations).Where("category_id = ?", uint(id)).Find(&menus).Error

	return menus, err
}

func CreateMenu(menu *models.Menu) error {
	return db.DB.Create(menu).Error
}

func UpdateMenu(menu *models.Menu) error {
	return db.DB.Save(menu).Error
}

func DeleteMenu(id int) error {
	var menu models.Menu
	return db.DB.Delete(&menu, id).Error
}
