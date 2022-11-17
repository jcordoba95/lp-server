package initializers

import (
	"log"
	"os"

	"github.com/jcordoba95/lp-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		DB.AutoMigrate(
			&models.User{},
			&models.Operation{},
			&models.Record{},
		)
		// TODO: Create operations with default costs if they don't exist already
	}
}
