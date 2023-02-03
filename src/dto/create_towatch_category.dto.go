package dto

type CreateTowatchCategoryDto struct {
	Value  string `json:"value"`
	Color  string `json:"color"`
	UserId int    `json:"user_id"`
}
