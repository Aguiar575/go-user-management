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

	var connectionString string = 
    fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, password, host, port)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
  
  createDatabase(databaseName, db)
  db.Close()

  connectionString = 
    fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, databaseName)
  database, err := sql.Open("mysql", connectionString)
  if err == nil {
    tableUp(database)
  }

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
  query := `
    CREATE TABLE users (
      id INT PRIMARY KEY AUTO_INCREMENT,
      username VARCHAR(255) NOT NULL UNIQUE,
      email VARCHAR(255) NOT NULL UNIQUE
    );
  `
  _, err := db.Exec(query)

  if err != nil {
    fmt.Println("Error applying migration:", err)
  } else {
    fmt.Println("Users table created successfully!")
  }
}
