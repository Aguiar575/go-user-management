package main

import (
	"os"

	config "user-management/internal/configuration"
	db "user-management/internal/database"
	services "user-management/internal/services"
)

func main() {
	config.StartEnv()

	database, err := db.ConfigureDatabase(os.Getenv("HOST"), os.Getenv("PORT"),
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	defer database.Close()

	PermissionsRepository := db.NewUsersPermissionsRepository(database)
	permissionsService := services.NewPermissionsService(PermissionsRepository)
}
