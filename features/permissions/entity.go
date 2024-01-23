package permissions

import (
	"github.com/labstack/echo/v4"
)

type Permissions struct {
	Code string
	Name string
}

type Handler interface {
	GetAllPermissions() echo.HandlerFunc
	AddPermissions() echo.HandlerFunc
	DeletePermissions() echo.HandlerFunc
}

type Repository interface {
	GetAllPermissions() ([]Permissions, error)
	AddPermissions(newPermission Permissions) (Permissions, error)
	DeletePermissions(code string) error
}

type Services interface {
	GetAllPermissions() ([]Permissions, error)
	AddPermissions(newPermission Permissions) (Permissions, error)
	DeletePermissions(code string) error
}
