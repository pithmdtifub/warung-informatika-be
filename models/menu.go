package models

type Menu struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	CategoryID  uint   `gorm:"not null"`
	Category    Category
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Image       string  `gorm:"not null"`
}
