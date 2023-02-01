package models

type TowatchCategory struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Value  string `json:"value"`
	Color  string `json:"color"`
	UserId int    `json:"user_id"`
}
