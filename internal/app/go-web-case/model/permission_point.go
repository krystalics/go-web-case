package model

type PermissionPoint struct {
	Model
	PermissionType string `json:"permission_type"`
	AppId          string `json:"app_id"`
}
