package model

import "time"

type RolePermission struct {
	Id             int64     `json:"id"`
	RoleId         int64     `json:"role_id"`
	PermissionId   int64     `json:"permission_id"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	Ctime          time.Time `json:"ctime"`
	Mtime          time.Time `json:"mtime"`
}
