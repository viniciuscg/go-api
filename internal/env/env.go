package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvString(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
