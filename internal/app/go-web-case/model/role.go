package model

import "time"

// Role omitempty 表示如果在序列化的时候、如果不存在则不进行默认值填充。结果就不包含这个字段
//参考 https://old-panda.com/2019/12/11/golang-omitempty/
type Role struct {
	Id             int64     `json:"id"`
	RoleName       string    `json:"role_name"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	Ctime          time.Time `json:"ctime"`
	Mtime          time.Time `json:"mtime"`
}
