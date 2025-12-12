package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetProperty(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env property")
		return ""
	}
	return os.Getenv(key)
}
