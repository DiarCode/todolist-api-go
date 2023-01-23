package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Todos     []Todo    `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"todos"`
	CreatedAt time.Time `json:"created_at"`
}
