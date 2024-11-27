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