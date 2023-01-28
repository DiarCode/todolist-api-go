package dto

import "time"

type CreateTowatchDto struct {
	Title      string    `json:"title"`
	StartDate  time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date"`
	Episodes   int       `json:"episodes"`
	Rating     float32   `json:"rating"`
	Studio     string    `json:"studio"`
	Image      string    `json:"image"`
}
