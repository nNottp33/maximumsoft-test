package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	JWT_SECRET string
	DB_URL     string
)

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	DB_URL = os.Getenv("DB_URL")

}
