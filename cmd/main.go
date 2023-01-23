package main

import (
	"log"

	"github.com/DiarCode/todo-go-api/pkg/config/database/postgres"
	"github.com/DiarCode/todo-go-api/pkg/config/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

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

	postgres.ConnectDB()

	log.Fatal(app.Listen(":8080"))
}
