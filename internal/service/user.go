package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type LogoutRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) Logout(param *LogoutRequest) {
	// todo: record logout event
}

type CreateUserRequest struct {
	Token    string `json:"token" form:"token"`
	Username string `json:"username" form:"username"`
}

func (s *Service) CreateUser(param *CreateUserRequest) uint {
	return s.dao.CreateUser(param.Token, param.Username)
}

type GetUserRequest struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) GetUser(param *GetUserRequest) *model.User {
	return s.dao.GetUser(param.Token)
}
