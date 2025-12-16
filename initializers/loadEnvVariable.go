package initializers

import (
	"github.com/lpernett/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}