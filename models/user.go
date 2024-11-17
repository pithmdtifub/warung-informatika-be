package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string
	Role     string `gorm:"type:role;not null;default:'Admin'"`
	Password string
}

const (
	RoleUser  = "User"
	RoleAdmin = "Admin"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	if u.Username == "" || u.Password == "" {
		err = errors.New("can't save invalid data")
	}
	return
}
