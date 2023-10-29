package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {
	envPath := "/Users/marcobraun/dev/Tekbit/server/.env"

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file at path: %s", envPath)
	}
}
