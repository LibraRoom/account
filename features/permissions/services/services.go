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
