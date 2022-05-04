package model

import "time"

type App struct {
	Id             int64     `json:"id"`     //自增主键
	AppId          string    `json:"app_id"` //appId
	AppName        string    `json:"app_name"`
	OrgId          string    `json:"org_id"`
	OrgName        string    `json:"org_name"`
	OwnerName      string    `json:"owner_name"`
	OwnerMail      string    `json:"owner_mail"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	Ctime          time.Time `json:"ctime"`
	Mtime          time.Time `json:"mtime"`
}
