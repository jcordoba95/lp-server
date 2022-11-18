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
		// TODO: Find a better way to populate initial DB instead of this + unique_index on Type. Create seeds for Operation.
		DB.Create([]models.Operation{
			{Type: "addition", Cost: 10},
			{Type: "subtraction", Cost: 20},
			{Type: "multiplication", Cost: 30},
			{Type: "division", Cost: 40},
			{Type: "square_root", Cost: 50},
			{Type: "random_string", Cost: 60},
		})
	}
}
