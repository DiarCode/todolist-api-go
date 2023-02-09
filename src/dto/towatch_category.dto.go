package dto

type CreateTowatchCategoryDto struct {
	Value  string `json:"value"`
	Color  string `json:"color"`
	UserId int    `json:"user_id"`
}

type AddTowatchToCategoryDto struct {
	UserId int    `json:"user_id"`
	TowatchCategoryId int `json:"towatch_category_id"`
	TowatchId int `json:"towatch_id"`
}