package db

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(host string, port string, username string, password string) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, password, host, port)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    return db, nil
}
