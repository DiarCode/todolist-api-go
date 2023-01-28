package routes

import (
	"log"

	"github.com/DiarCode/todo-go-api/src/controllers"
	"github.com/DiarCode/todo-go-api/src/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func InitRoutes(app *fiber.App) {
	router := app.Group("api/v1")

	todoRouter := router.Group("/todos")
	todoRouter.Get("/", middleware.AuthMiddleware(), controllers.GetAllTodos)
	todoRouter.Post("/", controllers.CreateTodo)
	todoRouter.Get("/:id", controllers.GetTodoById)
	todoRouter.Delete("/:id", controllers.DeleteTodoById)

	todoCategoryRouter := router.Group("/todos-category")
	todoCategoryRouter.Get("/", controllers.GetAllTodoCategories)
	todoCategoryRouter.Post("/", controllers.CreateTodoCategory)
	todoCategoryRouter.Get("/:id", controllers.GetTodoCategoryById)
	todoCategoryRouter.Delete("/:id", controllers.DeleteTodoCategoryById)

	towatchRouter := router.Group("/towatch")
	towatchRouter.Get("/", controllers.GetAllTowatch)
	towatchRouter.Post("/", controllers.CreateTowatch)
	towatchRouter.Get("/:id", controllers.GetTowatchById)
	towatchRouter.Delete("/:id", controllers.DeleteTowatchById)

	towatchCategoryRouter := router.Group("/towatch-category")
	towatchCategoryRouter.Get("/", controllers.GetAllTowatchCategories)
	towatchCategoryRouter.Post("/", controllers.CreateTowatchCategory)
	towatchCategoryRouter.Get("/:id", controllers.GetTowatchCategoryById)
	towatchCategoryRouter.Delete("/:id", controllers.DeleteTowatchCategoryById)

	userRouter := router.Group("/users")
	userRouter.Get("/", controllers.GetAllUsers)
	userRouter.Post("/", controllers.CreateUser)
	userRouter.Get("/:id", controllers.GetUserById)
	userRouter.Delete("/:id", controllers.DeleteUserById)

	authRouter := router.Group("/auth")
	authRouter.Post("/login", controllers.Login)
	authRouter.Post("/signup", controllers.Signup)

	userTowatchRouter := router.Group("/user-towatch")
	userTowatchRouter.Get("/", controllers.GetAllTowathesByCategory)
	userTowatchRouter.Post("/", controllers.AssignTowatchToCategory)

	log.Println("")

}
