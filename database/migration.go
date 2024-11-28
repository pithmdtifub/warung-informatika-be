package database

import (
	"log"
	"warung-informatika-be/models"
)

func MigrateUp() {
	DB.Exec(`DO $$ BEGIN CREATE TYPE role AS ENUM ('User', 'Admin'); EXCEPTION WHEN duplicate_object THEN null; END $$;`)

	err := DB.AutoMigrate(&models.Category{}, &models.Menu{}, &models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MigrateDown() {
	err := DB.Migrator().DropTable(&models.Category{}, &models.Menu{}, &models.User{})

	DB.Exec("DROP TYPE IF EXISTS role;")
	if err != nil {
		log.Fatal(err.Error())
	}
}
