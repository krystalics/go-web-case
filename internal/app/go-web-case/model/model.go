package model

import "time"

//gorm支持将各表中的公共字段抽出来、在具体的实体中嵌入Model即可
type Model struct {
	Id             int64     `json:"id"` //自增主键
	IsDeleted      bool      `json:"is_deleted"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	Ctime          time.Time `json:"ctime"`
	Mtime          time.Time `json:"mtime"`
}
