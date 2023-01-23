package routes

import (
	"github.com/DiarCode/todo-go-api/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func InitRoutes(app *fiber.App) {
	router := app.Group("api/v1")

	todoRouter := router.Group("/todos")
	todoRouter.Get("/", controllers.GetAllTodos)
	todoRouter.Post("/", controllers.CreateTodo)
	todoRouter.Get("/:id", controllers.GetTodoById)
	todoRouter.Delete("/:id", controllers.DeleteTodoById)

	userRouter := router.Group("/users")
	userRouter.Get("/", controllers.GetAllUsers)
	todoRouter.Post("/", controllers.CreateUser)
	todoRouter.Get("/:id", controllers.GetUserById)
	todoRouter.Delete("/:id", controllers.DeleteUserById)

}
