package db

import (
	"database/sql"
  models "user-management/internal/models"
)

type UsersPermissionsRepository interface {
  AddPermission(permission *models.UserPermission) error
  GetUserPermissions(userID int) ([]models.UserPermission, error)
}

type usersPermissionsRepositoryRepository struct {
  db *sql.DB
}

func NewUsersPermissionsRepository(db *sql.DB) UsersPermissionsRepository {
  return &usersPermissionsRepositoryRepository{db: db}
}

func (ur *usersPermissionsRepositoryRepository) AddPermission(permission *models.UserPermission) error {
	query := `
		INSERT INTO user_permissions (user_id, context_id, read, write)
		VALUES (?, ?, ?, ?)
	`
	_, err := ur.db.Exec(query, permission.UserID, permission.ContextID, permission.Read, permission.Write)
	return err
}

func (ur *usersPermissionsRepositoryRepository) GetUserPermissions(userID int) ([]models.UserPermission, error) {
	var permissions []models.UserPermission

	rows, err := ur.db.Query("SELECT * FROM user_permissions WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission models.UserPermission
		err := rows.Scan(&permission.ID, &permission.UserID, &permission.ContextID, &permission.Read, &permission.Write)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
