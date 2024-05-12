package db

import (
	"database/sql"
  models "user-management/internal/models"
)

type UserRepository interface {
  RegisterUser(user *models.User) error
  GetUserById(id int) (*models.User, error)
  DeleteUser(id int) error
}

type userRepository struct {
  db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
  return &userRepository{db: db}
}

func (ur *userRepository) RegisterUser(user *models.User) error {
  query := "INSERT INTO users (username, email) VALUES (?, ?)"
  _, err := ur.db.Exec(query, user.Username, user.Email)
  return err
}

func (ur *userRepository) GetUserById(id int) (*models.User, error) {
  var user models.User

  rows, err := ur.db.Query("SELECT * FROM users WHERE id = ?", id)
  if err != nil {
    return nil, err
  }
  defer rows.Close() // Close the rows after use

  if rows.Next() {
    err := rows.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
      return nil, err
    }
    return &user, nil
  }

  return nil, nil // User not found
}

func (ur *userRepository) DeleteUser(id int) error {
  query := "DELETE FROM users WHERE id = ?"
  _, err := ur.db.Exec(query, id)
  return err
}
