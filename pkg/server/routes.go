package server

import (
	"github.com/DiarCode/todo-go-api/pkg/user"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	// User
	router.GET("/api/v1/users", user.GetAllUsers)
	router.POST("/api/v1/users/", user.AddUser)
	router.GET("/api/v1/users/:id", user.GetUserById)

	return router
}
