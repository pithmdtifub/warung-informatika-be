package models

type Menu struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	Name         string   `json:"name" form:"name" validate:"required"`
	CategoryID   uint     `json:"category_id" form:"category_id" validate:"required,number"`
	Category     Category `json:"-"`
	CategoryName string   `json:"category" gorm:"->"`
	Description  string   `json:"description" form:"description" validate:"required"`
	Price        float64  `json:"price" form:"price" validate:"required,number"`
	Image        string   `json:"image" form:"image" validate:"required,url"`
}
