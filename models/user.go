package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username" validate:"required"`
	Role     string `json:"role" gorm:"type:role;not null;default:'Admin'"`
	Password string `json:"password" form:"password" validate:"required"`
}
