package model

import "time"

type PermissionPoint struct {
	Id             int64     `json:"id"`
	PermissionType string    `json:"permission_type"`
	AppId          string    `json:"app_id"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	Ctime          time.Time `json:"ctime"`
	Mtime          time.Time `json:"mtime"`
}
