package models

type Menu struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	Name         string   `json:"name" gorm:"column:name"`
	CategoryID   uint     `json:"category_id" gorm:"column:category_id"`
	Category     Category `json:"-"`
	CategoryName string   `json:"category" gorm:"->"`
	Description  string   `json:"description" gorm:"column:description"`
	Price        float64  `json:"price" gorm:"column:price"`
	Image        string   `json:"image" gorm:"column:image"`
}
