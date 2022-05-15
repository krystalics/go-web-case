package service

import (
	"go-web-case/internal/app/go-web-case/conf"
	"go-web-case/internal/app/go-web-case/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

var srv *Service

func New(c *conf.Config) (s *Service) {
	srv = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}
