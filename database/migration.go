package database

import (
	"log"
	"warung-informatika-be/models"
)

func Migrate() {
	err := DB.Migrator().DropTable(&models.Category{}, &models.Menu{}, &models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = DB.AutoMigrate(&models.Category{}, &models.Menu{}, &models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
