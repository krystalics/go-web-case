package dao

import (
	"go-web-case/internal/app/go-web-case/model"
)

func (d *Dao) CreateUser(user model.User) (bool, error) {
	tx := d.db.Create(&user)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
