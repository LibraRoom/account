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
				"message": "Failed to add Permissions Duplicate Input",
			})
		}

		var response = new(PermissionsResponse)
		response.Code = result.Code
		response.Name = result.Name
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "Success add new permission",
			"data":    result,
		})
	}
}

// DeletePermissions implements permissions.Handler.
func (ph *PermissionsHandler) DeletePermissions() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")
		if code == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid permission code",
				"data":    nil,
			})
		}

		errDel := ph.p.DeletePermissions(code)

		if errDel != nil {
			c.Logger().Error("ERROR Deleting Permissions, explain:", errDel.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			return c.JSON(statusCode, map[string]interface{}{
				"message": message,
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Delete Permissions Success",
		})
	}
}

// UpdatePermissions implements permissions.Handler.
func (ph *PermissionsHandler) UpdatePermissions() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputData = new(PermissionsPutRequest)
		code := c.Param("code")
		if code == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid permission code",
				"data":    nil,
			})
		}

		if err := c.Bind(inputData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
				"data":    nil,
			})
		}

		updatePermission := permissions.Permissions{
			Name: inputData.Name,
		}

		result, err := ph.p.UpdatePermissions(code, updatePermission)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "Failed to update Permissions Duplicate Input",
			})
		}
		var response = new(PermissionsResponse)
		response.Code = result.Code
		response.Name = result.Name
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "Success Update permission Data",
			"data":    result,
		})
	}
}
