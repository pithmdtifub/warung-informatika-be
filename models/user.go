package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
