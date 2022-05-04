package model

//用户一般是从uc那边获取、权限系统的重点不在用户，也不需要创建
type User struct {
	Ucid      int64  `json:"ucid"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	IsDeleted bool   `json:"is_deleted"`
}
