package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Todos     []Todo    `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"todos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
