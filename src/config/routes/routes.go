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

	todoCategoryRouter := router.Group("/todos-category")
	todoCategoryRouter.Get("/", controllers.GetAllTodoCategories)
	todoCategoryRouter.Post("/", controllers.CreateTodoCategory)
	todoCategoryRouter.Get("/:id", controllers.GetTodoCategoryById)
	todoCategoryRouter.Delete("/:id", controllers.DeleteTodoCategoryById)

	userRouter := router.Group("/users")
	userRouter.Get("/", controllers.GetAllUsers)
	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/:id", controllers.GetUserById)
	userRouter.Delete("/:id", controllers.DeleteUserById)

	authRouter := router.Group("/auth")
	authRouter.Post("/login", controllers.Login)
	authRouter.Post("/signup", controllers.Signup)

}
