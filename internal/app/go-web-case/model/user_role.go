package model

type UserRole struct {
	Model
	Ucid   int64 `json:"ucid"`
	RoleId int64 `json:"role_id"`
}
