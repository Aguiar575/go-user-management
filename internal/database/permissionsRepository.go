package db

import (
	"database/sql"
	models "user-management/internal/models"
)

type PermissionsRepository interface {
	AddPermission(permission *models.UserPermission) error
	GetUserPermissions(userID int) ([]models.UserPermission, error)
	DeletePermission(permissionID int) error
}

type permissionsRepositoryImpl struct {
	db *sql.DB
}

func NewUsersPermissionsRepository(db *sql.DB) PermissionsRepository {
	return &permissionsRepositoryImpl{db: db}
}

func (ur *permissionsRepositoryImpl) AddPermission(permission *models.UserPermission) error {
	query := `
		INSERT INTO user_permissions (user_id, context_id, read, write)
		VALUES (?, ?, ?, ?)
	`
	_, err := ur.db.Exec(query, permission.UserID, permission.ContextID, permission.Read, permission.Write)
	return err
}

func (ur *permissionsRepositoryImpl) GetUserPermissions(userID int) ([]models.UserPermission, error) {
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

func (ur *permissionsRepositoryImpl) DeletePermission(permissionID int) error {
	query := "DELETE FROM user_permissions WHERE id = ?"
	_, err := ur.db.Exec(query, permissionID)
	return err
}
