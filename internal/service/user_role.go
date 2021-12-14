package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type CreateUserRoleRequest struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	RoleName string `json:"role_name" form:"role_name"`
}

func (s *Service) CreateUserRole(param *CreateUserRoleRequest) uint {
	return s.dao.CreateUserRole(param.UserID, param.RoleName)
}

type GetUserRoleRequest struct {
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *Service) GetUserRoles(param *GetUserRoleRequest) []*model.UserRole {
	return s.dao.GetUserRoles(param.UserID)
}
