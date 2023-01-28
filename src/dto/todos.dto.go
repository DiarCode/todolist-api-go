package dto

type CreateTodoDto struct {
	Title      string `json:"title"`
	Priority   bool   `json:"priority"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"category_id"`
}

type TodoByCategoryDto struct {
	UserId     int `json:"user_id"`
	CategoryId int `json:"category_id"`
}
