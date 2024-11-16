package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/models"

	"gorm.io/gorm/clause"
)

func GetMenus(limit, offset int) ([]models.Menu, error) {
	var menus []models.Menu
	err := db.DB.Preload(clause.Associations).Limit(limit).Offset(offset).Find(&menus).Error

	for i := range menus {
		menus[i].CategoryName = menus[i].Category.Name
	}

	return menus, err
}

func GetMenu(id int) (models.Menu, error) {
	var menu models.Menu
	err := db.DB.Preload(clause.Associations).First(&menu, id).Error

	menu.CategoryName = menu.Category.Name

	return menu, err
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
