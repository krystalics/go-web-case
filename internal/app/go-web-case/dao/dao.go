package dao

import (
	"context"
	"database/sql"
	"go-web-case/internal/app/go-web-case/conf"
	"gorm.io/gorm"
)

type Dao struct {
	conf *conf.Config
	db   *gorm.DB
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		conf: c,
		db:   conf.InitDB(c),
	}
	return
}

func (d *Dao) BeginTran(ctx context.Context) (tx *sql.Tx, err error) {
	return
}
