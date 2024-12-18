package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"not null;unique"`
	Role     string    `gorm:"type:role;not null;default:'Admin'"`
	Password string    `gorm:"not null"`
}

const (
	RoleUser  = "User"
	RoleAdmin = "Admin"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
