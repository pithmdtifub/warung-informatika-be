package repositories

import (
	db "warung-informatika-be/database"
	"warung-informatika-be/dto"
	"warung-informatika-be/helpers"
	"warung-informatika-be/models"

	"gorm.io/gorm/clause"
)

func GetMenus(query dto.MenuQuery) ([]models.Menu, error) {
	var menus []models.Menu
	queryDB := db.DB.Preload(clause.Associations)

	if query.Search != "" {
		queryDB = queryDB.Where("name ILIKE ?", "%"+query.Search+"%")
	}

	if query.Category != 0 {
		queryDB = queryDB.Where("category_id = ?", query.Category)
	}

	offset := (query.Page - 1) * query.Limit
	queryDB = queryDB.Limit(query.Limit).Offset(offset)

	err := queryDB.Find(&menus).Error

	return menus, err
}

func GetMenu(id uint) (models.Menu, error) {
	menu := models.Menu{ID: id}
	err := db.DB.Preload(clause.Associations).First(&menu).Error

	return menu, err
}

func CreateMenu(menu *models.Menu) error {
	return db.DB.Create(menu).Error
}

func UpdateMenu(menu *models.Menu) error {
	res := db.DB.Save(menu)

	if err := helpers.CheckRowsAffected(res.RowsAffected); err != nil {
		return err
	}

	return res.Error
}

func DeleteMenu(id int) error {
	var menu models.Menu
	res := db.DB.Delete(&menu, id)

	if err := helpers.CheckRowsAffected(res.RowsAffected); err != nil {
		return err
	}

	return res.Error
}
