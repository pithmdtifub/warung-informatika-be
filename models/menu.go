package models

type Menu struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	CategoryID  uint
	Category    Category
	Description string
	Price       float64
	Image       string
}
