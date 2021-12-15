package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type CreateUserRolesRequest struct {
	UserID    uint     `json:"user_id" form:"user_id"`
	RoleNames []string `json:"roles" form:"roles"`
}

func (s *Service) CreateUserRoles(param *CreateUserRolesRequest) []*model.UserRole {
	return s.dao.CreateUserRoles(param.UserID, param.RoleNames)
}

type GetUserRoleRequest struct {
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *Service) GetUserRoles(param *GetUserRoleRequest) []*model.UserRole {
	return s.dao.GetUserRoles(param.UserID)
}

type DeleteUserRolesRequest struct {
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *Service) DeleteUserRoles(param *DeleteUserRolesRequest) {
	s.dao.DeleteUserRoles(param.UserID)
}
