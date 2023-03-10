package main

import (
	"log"
	"os"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// 		POSTGRES_DSN = "host=localhost port=5432 user=postgres password=password dbname=gotodo_db sslmode=disable"

// 		# Надо сделать перед запуском
// 		# 1. Создать .env файл в корневой директории (там где go.mod)
// 		# 2. Засунуть туда POSTGRES_DSN, но изменить на свои парамметры (скопировать прям как сверху)

// 		# Запуск проекта
// 		# go run ./src

// 		# Пример endpoint
// 		# http://localhost:8080/api/v1/users

func main() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// environmentPath := filepath.Join(dir, ".env")
	// err = godotenv.Load(environmentPath)
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.InitRoutes(app)

	database.ConnectDB()

	log.Fatal(app.Listen(getPort()))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	return port
}
