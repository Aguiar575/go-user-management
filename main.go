package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	config "user-management/internal/configuration"
	db "user-management/internal/database"
	"user-management/internal/models"
	services "user-management/internal/services"
)

var permissionsService *services.PermissionsService

func main() {
	config.StartEnv()

	database, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to configure database: %v", err)
	}
	defer database.Close()

	setupServices(database)

	router := mux.NewRouter()
	router.HandleFunc("/permissions/{userID}", getPermissionsHandler).Methods("GET")
	router.HandleFunc("/permissions", setPermissionHandler).Methods("POST")
	router.HandleFunc("/permissions/{permissionID}", removePermissionHandler).Methods("DELETE")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

func setupDatabase() (*sql.DB, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	database, err := db.ConfigureDatabase(host, port, username, password)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func setupServices(database *sql.DB) {
	permissionsRepository := db.NewUsersPermissionsRepository(database)
	permissionsService = services.NewPermissionsService(permissionsRepository)
}

func getPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	permissions, err := permissionsService.GetPermissionsByUserId(userID)
	if err != nil {
		http.Error(w, "Error fetching permissions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func setPermissionHandler(w http.ResponseWriter, r *http.Request) {
	var permission models.UserPermission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := permissionsService.SetPermission(&permission); err != nil {
		http.Error(w, "Error setting permission", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func removePermissionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	permissionID, err := strconv.Atoi(vars["permissionID"])
	if err != nil {
		http.Error(w, "Invalid permission ID", http.StatusBadRequest)
		return
	}

	if err := permissionsService.RemovePermission(permissionID); err != nil {
		http.Error(w, "Error removing permission", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
