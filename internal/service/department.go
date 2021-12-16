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

type UpdateDepartmentRequest struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name" `
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) UpdateDepartment(param *UpdateDepartmentRequest) (*model.Department, error) {
	return s.dao.UpdateDepartment(param.ID, param.Name, param.Introduction)
}

type DeleteDepartmentRequest struct {
	ID uint `json:"id" form:"id"`
}

func (s *Service) DeleteDepartment(param *DeleteDepartmentRequest) error {
	err := s.dao.DeleteDepartment(param.ID)
	if err != nil {
		return err
	}
	s.dao.DeleteDepartmentUsers(param.ID)
	return nil
}
