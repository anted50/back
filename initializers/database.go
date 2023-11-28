package initializers

import (
	"test/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBGorm *gorm.DB

func InitGorm() error {
	psqlInfo := "host=localhost user=postgres password=1234 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Trip{})
	DBGorm = db
	return nil
}
