package routes

import (
	"account/features/permissions"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, ph permissions.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routePermissions(e, ph)
}

func routePermissions(e *echo.Echo, ph permissions.Handler) {
	e.GET("/permissions", ph.GetAllPermissions())
}
