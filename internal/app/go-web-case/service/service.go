package service

import (
	"go-web-case/internal/app/go-web-case/conf"
	"go-web-case/internal/app/go-web-case/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}
