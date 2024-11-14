package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Role     string `json:"role" form:"role"`
	Password string `json:"-"`
}
