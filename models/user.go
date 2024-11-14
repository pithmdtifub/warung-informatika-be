package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Role     string `json:"role" gorm:"type:role;not null;default:'Admin'"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
