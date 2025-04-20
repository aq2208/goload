package configs

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found, using system environment variables")
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return ""
}