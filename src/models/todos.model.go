package models

type Todo struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserId      int    `json:"-"`
}
