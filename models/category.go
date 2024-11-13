package models

type Category struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" form:"name"`
	Menus []Menu `json:"-"`
}
