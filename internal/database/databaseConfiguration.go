package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConfigureDatabase(
	host string, port string,
	username string, password string) (*sql.DB, error) {
	databaseName := "usermanagement"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, password, host, port)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	createDatabase(databaseName, db)
	defer db.Close() // Close the connection after table creation (if successful)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, databaseName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	tableUp(database) // Ensure tables are created before returning

	return database, err
}

func createDatabase(databaseName string, db *sql.DB) {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", databaseName)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func tableUp(db *sql.DB) {
	userTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY AUTO_INCREMENT,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE
		);
	`
	_, err := db.Exec(userTableQuery)
	if err != nil {
		fmt.Println("Error creating users table:", err)
		return
	}

	permissionTableQuery := `
		CREATE TABLE IF NOT EXISTS user_permissions (
			id INT PRIMARY KEY AUTO_INCREMENT,
			user_id INT NOT NULL,
			context_id INT NOT NULL,
			read BOOL NOT NULL,
			write BOOL NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	_, err = db.Exec(permissionTableQuery)
	if err != nil {
		fmt.Println("Error creating user_permissions table:", err)
	} else {
		fmt.Println("Tables created successfully!")
	}
}

