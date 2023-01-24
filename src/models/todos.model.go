package models

import "time"

type Todo struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Completed  bool      `json:"completed"`
	UserId     int       `json:"-"`
	Priority   bool      `json:"priority"`
	CreatedAt  time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	CategoryId int       `gorm:"foreignKey:CategoryId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"category"`
}
