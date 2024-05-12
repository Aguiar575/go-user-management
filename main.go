package main

import (
  "database/sql"
  "fmt"
	"os"

  models "user-management/internal/models"
  db "user-management/internal/database"
  config "user-management/internal/configuration"
  services "user-management/internal/services"
)

func main() {
  config.StartEnv()

	database, err := db.ConnectToDB(os.Getenv("HOST"), os.Getenv("PORT"),
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}  
  defer database.Close() // Close the connection when done

  // Create user repository and service
  userRepository := db.NewUserRepository(database)
  userService := services.NewUserService(userRepository)

  // Create a new user
  newUser := &models.User{
    Username: "janedoe",
    Email:    "janedoe@example.com",
  }

  // Register the user
  err = userService.RegisterUser(newUser)
  if err != nil {
    fmt.Println("Error registering user:", err)
    return
  }

  fmt.Println("User registered successfully!")

  // Retrieve user by ID
  user, err := userRepository.GetUserById(newUser.ID)
  if err != nil {
    if err == sql.ErrNoRows {
      fmt.Println("User not found")
    } else {
      fmt.Println("Error retrieving user:", err)
    }
    return
  }
  fmt.Printf("Retrieved user: %+v\n", user)

  // Delete user by ID (assuming proper authorization)
  err = userRepository.DeleteUser(newUser.ID)
  if err != nil {
    fmt.Println("Error deleting user:", err)
    return
  }
  fmt.Println("User deleted successfully")
}

