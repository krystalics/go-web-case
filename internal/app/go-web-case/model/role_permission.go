package model

type RolePermission struct {
	Model
	RoleId       int64 `json:"role_id"`
	PermissionId int64 `json:"permission_id"`
}
