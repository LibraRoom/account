package services

import (
	"account/features/permissions"
)

type PermissionsServices struct {
	p permissions.Repository
}

func New(model permissions.Repository) permissions.Services {
	return &PermissionsServices{
		p: model,
	}
}

// GetAllPermissions implements permissions.Services.
func (ps *PermissionsServices) GetAllPermissions() ([]permissions.Permissions, error) {
	result, err := ps.p.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AddPermissions implements permissions.Services.
func (ps *PermissionsServices) AddPermissions(newPermission permissions.Permissions) (permissions.Permissions, error) {
	result, err := ps.p.AddPermissions(newPermission)
	if err != nil {
		return permissions.Permissions{}, err
	}

	return result, nil
}
