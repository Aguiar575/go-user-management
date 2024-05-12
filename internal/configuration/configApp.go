package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func StartEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

