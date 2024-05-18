package services

import (
	db "user-management/internal/database"
	"user-management/internal/models"
)

type PermissionsService struct {
	permissionsRepository db.PermissionsRepository
}

func NewPermissionsService(permissionsRepository db.PermissionsRepository) *PermissionsService {
	return &PermissionsService{permissionsRepository: permissionsRepository}
}

func (ps *PermissionsService) GetPermissionsByUserId(userID int) ([]models.UserPermission, error) {
	return ps.permissionsRepository.GetUserPermissions(userID)
}

func (ps *PermissionsService) SetPermission(permission *models.UserPermission) error {
	return ps.permissionsRepository.AddPermission(permission)
}

func (ps *PermissionsService) RemovePermission(permissionID int) error {
	return ps.permissionsRepository.DeletePermission(permissionID)
}
