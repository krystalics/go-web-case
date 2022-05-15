package service

import "go-web-case/internal/app/go-web-case/model"

func (s *Service) CreateUser(user model.User) (bool, error) {
	return s.dao.CreateUser(user)
}
