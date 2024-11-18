package dto

type MenuRequest struct {
	Name        string  `json:"name" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Image       string  `json:"image" validate:"required,url"`
}

type MenuUpdateRequest struct {
	Name        string  `json:"name" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Image       string  `json:"image" validate:"required,url"`
}

type MenuResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
}
