package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username" validate:"required"`
	Role     string `json:"role"`
	Password string `json:"-" form:"password" validate:"required"`
}
