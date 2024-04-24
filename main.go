package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	db "user-management/internal"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

	db, err := db.ConnectToDB(os.Getenv("HOST"), os.Getenv("PORT"),
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Hello, World!")
}
