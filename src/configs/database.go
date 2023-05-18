package configs

import (
	"github.com/nNottp33/maximumsoft-test/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func ConnectDatabase() {
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Employees{})

	Db = db
}
