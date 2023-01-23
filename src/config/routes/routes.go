package routes

import (
	"github.com/DiarCode/todo-go-api/src/controllers"
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
	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/:id", controllers.GetUserById)
	userRouter.Delete("/:id", controllers.DeleteUserById)

}
