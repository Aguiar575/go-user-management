package config

import (
	"github.com/joho/godotenv"
	"fmt"
)

func StartEnv() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
}
