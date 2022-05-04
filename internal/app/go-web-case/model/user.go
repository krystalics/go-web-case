package model

type User struct {
	Ucid      int64  `json:"ucid"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	IsDeleted bool   `json:"is_deleted"`
}
