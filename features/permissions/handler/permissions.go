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

// AddPermissions implements permissions.Handler.
func (ph *PermissionsHandler) AddPermissions() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputData = new(PermissionsRequest)
		if err := c.Bind(&inputData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(permissions.Permissions)
		inputProcess.Code = inputData.Code
		inputProcess.Name = inputData.Name

		result, err := ph.p.AddPermissions(*inputProcess)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"massage": "Failed to add Permissions",
			})
		}

		var response = new(PermissionsResponse)
		response.Code = result.Code
		response.Name = result.Name
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "Success fetching all permission data",
			"data":    result,
		})
	}
}
