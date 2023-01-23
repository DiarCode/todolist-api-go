package postgres

import (
	"log"
	"os"

	"github.com/DiarCode/todo-go-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = os.Getenv("POSTGRES_DSN")
)

func ConnectDB() {
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.Todo{})

	if err != nil {
		log.Fatal(err)
	}
}
