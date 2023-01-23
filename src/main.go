package main

import (
	"log"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/config/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// 		POSTGRES_DSN = "host=localhost port=5432 user=postgres password=password dbname=gotodo_db sslmode=disable"

// 		# Надо сделать перед запуском
// 		# 1. Создать .env файл в корнивой директории (там где go.mod)
// 		# 2. Засунуть туда POSTGRES_DSN, но изменить на свои парамметры (скопировать прям как сверху)

// 		# Запуск проекта
// 		# go run ./src

// 		# Пример endpoint
// 		# http://localhost:8080/api/v1/users

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.InitRoutes(app)

	database.ConnectDB()

	log.Fatal(app.Listen(":8080"))
}
