package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type GetRoleRequest struct {
	Name uint `json:"name" form:"name"`
}

func (s *Service) ListRoles() []*model.Role {
	return s.dao.ListRoles()
}
