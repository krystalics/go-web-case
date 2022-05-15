package dao

import (
	"go-web-case/internal/app/go-web-case/model"
)

func (d *Dao) CreateUser(user model.User) bool {
	d.db.Create(&user)
	return true
}
