package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type CreateUserDepartmentsRequest struct {
	UserID        uint   `json:"user_id" form:"user_id"`
	DepartmentIDs []uint `json:"departments" form:"departments"`
}

func (s *Service) CreateUserDepartments(param *CreateUserDepartmentsRequest) []*model.UserDepartment {
	return s.dao.CreateUserDepartments(param.UserID, param.DepartmentIDs)
}

type GetUserDepartmentsRequest struct {
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *Service) GetUserDepartments(param *GetUserDepartmentsRequest) []*model.UserDepartment {
	return s.dao.GetUserDepartments(param.UserID)
}

type DeleteUserDepartmentsRequest struct {
	UserID uint `json:"user_id" form:"user_id"`
}

func (s *Service) DeleteUserDepartments(param *DeleteUserDepartmentsRequest) {
	s.dao.DeleteUserDepartments(param.UserID)
}
