package models

type TodoCategory struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Value string `json:"value"`
	Color string `json:"color"`
}
