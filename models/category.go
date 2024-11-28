package models

import "gorm.io/gorm"

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Menus []Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Category) BeforeFind(tx *gorm.DB) (err error) {
	tx.Order("id ASC`")
	return
}
