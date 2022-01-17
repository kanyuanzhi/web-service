package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (s *Service) ListRoles() ([]*model.Role, error) {
	return s.dao.ListRoles()
}

type UpdateRoleRequest struct {
	ID           uint   `json:"id" form:"id"`
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) UpdateRole(param *UpdateRoleRequest) (*model.Role, error) {
	return s.dao.UpdateRole(param.ID, param.Introduction)
}
