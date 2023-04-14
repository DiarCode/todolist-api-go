package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockAppStruct struct {
	app *fiber.App
	ctx *fiber.Ctx
}

var MockApp = &MockAppStruct{}

func TestMain(m *testing.M) {
	//Before all
	connectMockDatabase()
	setupRoutes()

	code := m.Run()
	//After all
	clearMockDatabase()
	os.Exit(code)
}

func connectMockDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_PORT"),
	)
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	log.Println("Connected to database!")
	log.Println("Running migrations")

	err = database.DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.TodoCategory{},
		&models.Towatch{},
		&models.TowatchCategory{},
	)

	if err != nil {
		log.Fatal("Failed to migrate! \n", err)
	}

	log.Println("All migrations done!")
}

func clearMockDatabase() {
	log.Println("Clearing mock database!")

	database.DB.Migrator().DropTable(
		&models.User{},
		&models.Todo{},
		&models.TodoCategory{},
		&models.Towatch{},
		&models.TowatchCategory{},
	)

	log.Println("Mock database cleared!")
}

func setupRoutes() {
	app := fiber.New()
	routes.InitRoutes(app)
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	MockApp.app = app
	MockApp.ctx = c
}
