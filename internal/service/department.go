package service

import "github.com/kanyuanzhi/web-service/internal/model"

type CreateDepartmentRequest struct {
	Name         string `json:"name" form:"name" `
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) CreateDepartment(param *CreateDepartmentRequest) *model.Department {
	return s.dao.CreateDepartment(param.Name, param.Introduction)
}

type ListDepartmentsRequest struct{}

func (s *Service) ListDepartments() []*model.Department {
	return s.dao.ListDepartments()
}
