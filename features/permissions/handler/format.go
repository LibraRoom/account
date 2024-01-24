package handler

type PermissionsRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type PermissionsPutRequest struct {
	Name string `json:"name"`
}

type PermissionsResponse struct {
	Code string `json:"code" bson:"_id"`
	Name string `json:"name"`
}
