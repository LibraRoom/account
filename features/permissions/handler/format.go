package handler

type PermissionsRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type PermissionsResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
