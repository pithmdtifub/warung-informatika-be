package dto

type MenuParams struct {
	ID uint `params:"id"`
}

type MenuQuery struct {
	Search   string `query:"search"`
	Category int    `query:"category"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

type MenuDTO struct {
	Name         string  `json:"name" validate:"required"`
	CategoryID   uint    `json:"category_id" validate:"required,number"`
	CategoryName string  `json:"category_name"`
	Description  string  `json:"description" validate:"required"`
	Price        float64 `json:"price" validate:"required,number"`
	Image        string  `json:"image" validate:"required,url"`
}
