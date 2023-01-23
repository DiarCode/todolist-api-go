package dto

type CreateTodoDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}
