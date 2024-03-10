package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("example.env")
	if err != nil {
		log.Fatal("Error loading example.env file")
	}
}
