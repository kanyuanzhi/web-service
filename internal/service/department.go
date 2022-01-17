package service

import "github.com/kanyuanzhi/web-service/internal/model"

type CreateDepartmentRequest struct {
	Name         string `json:"name" form:"name" `
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) CreateDepartment(param *CreateDepartmentRequest) (*model.Department, error) {
	return s.dao.CreateDepartment(param.Name, param.Introduction)
}

type ListDepartmentsRequest struct{}

func (s *Service) ListDepartments() ([]*model.Department, error) {
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
	// 删除Departments表中的记录
	err := s.dao.DeleteDepartment(param.ID)
	if err != nil {
		return err
	}

	// 删除UserDepartmentAssociations表中的记录
	err = s.dao.DeleteUserDepartmentAssociationsByDepartmentID(param.ID)
	if err != nil {
		return err
	}

	return nil
}
