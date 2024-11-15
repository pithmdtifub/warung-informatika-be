package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Username string    `json:"username" form:"username" validate:"required"`
	Role     string    `json:"role" gorm:"type:role;not null;default:'Admin'"`
	Password string    `json:"password" form:"password" validate:"required"`
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
