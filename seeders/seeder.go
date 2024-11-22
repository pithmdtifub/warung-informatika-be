package seeders

import (
	"warung-informatika-be/models"
)

func Seed() {
	UserSeeder(models.User{Username: "admin", Password: "akusayangpit123"})
	UserSeeder(models.User{Username: "user", Role: models.RoleUser, Password: "akusayangpit123"})
}
