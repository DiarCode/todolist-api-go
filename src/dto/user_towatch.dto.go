package dto

type UserTowatchDto struct {
	CategoryID int `json:"category_id"`
	UserID     int `json:"user_id"`
}

type AssignTowatchToCategoryDto struct {
	TowatchID  int `json:"towatch_id"`
	UserID     int `json:"user_id"`
	CategoryID int `json:"category_id"`
}

type RemoveTowatchFromCategoryDto struct {
	TowatchID  int `json:"towatch_id"`
	UserID     int `json:"user_id"`
	CategoryID int `json:"category_id"`
}
