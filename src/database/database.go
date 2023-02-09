package database

import (
	"log"
	"os"

	"github.com/DiarCode/todo-go-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DATABASE_DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	log.Println("Connected to database!")
	log.Println("Running migrations")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.TodoCategory{},
		&models.Towatch{},
		&models.TowatchCategory{},
	)

	if err != nil {
		log.Fatal("Failed to connect to migrate! \n", err)
	}

	log.Println("Migrations done!")
}
