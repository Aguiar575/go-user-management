package main

import (
	"fmt"
	"os"
	db "user-management/internal/database"
  config "user-management/internal/configuration"
)

func main() {
  config.StartEnv()

	db, err := db.ConnectToDB(os.Getenv("HOST"), os.Getenv("PORT"),
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Hello, World!")
}
