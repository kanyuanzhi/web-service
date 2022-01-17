package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateDepartment(name string, introduction string) (*model.Department, error) {
	department := &model.Department{
		Name:         name,
		Introduction: introduction,
	}
	return department.Create()
}

func (dao *Dao) ListDepartments() ([]*model.Department, error) {
	department := &model.Department{}
	return department.List()
}

func (dao *Dao) UpdateDepartment(id uint, name string, introduction string) (*model.Department, error) {
	department := &model.Department{
		DefaultFields: &model.DefaultFields{ID: id},
		Name:          name,
		Introduction:  introduction,
	}
	return department.Update()
}

func (dao *Dao) DeleteDepartment(id uint) error {
	department := &model.Department{
		DefaultFields: &model.DefaultFields{ID: id},
	}
	return department.Delete()
}
