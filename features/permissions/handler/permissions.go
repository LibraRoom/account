package handler

import (
	"account/features/permissions"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type PermissionsHandler struct {
	p permissions.Services
}

func New(p permissions.Services) permissions.Handler {
	return &PermissionsHandler{
		p: p,
	}
}

// GetAllPermissions implements permissions.Handler.
func (ph *PermissionsHandler) GetAllPermissions() echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := ph.p.GetAllPermissions()
		if err != nil {
			c.Logger().Error("Error fetching permissions: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to retrieve permission data",
			})
		}
		var response []PermissionsResponse
		for _, result := range results {
			response = append(response, PermissionsResponse{
				Code: result.Code,
				Name: result.Name,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success fetching all permission data",
			"data":    response,
		})
	}
}
