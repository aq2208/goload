package configs

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("No .env file found, using system environment variables")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}