package models

import "time"

type Towatch struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	StartDate  time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date"`
	Episodes   int       `json:"episodes"`
	Rating     float32   `json:"rating"`
	Studio     string    `json:"studio"`
	Image      string    `json:"image"`
}
