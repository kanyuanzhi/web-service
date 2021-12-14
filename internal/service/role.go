package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type GetRoleRequest struct {
	Name uint `json:"name" form:"name"`
}

func (s *Service) GetRole(param *GetRoleRequest) *model.Role {
	return s.dao.GetRole(param.Name)
}
