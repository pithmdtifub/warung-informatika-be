package models

type Category struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" form:"name" validate:"required"`
	Menus []Menu `json:"-"`
}
