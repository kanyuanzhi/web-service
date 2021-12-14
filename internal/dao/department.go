package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateDepartment(name string, introduction string) *model.Department {
	department := &model.Department{
		Name:         name,
		Introduction: introduction,
	}
	return department.Create()
}

func (dao *Dao) ListDepartments() []*model.Department {
	department := &model.Department{}
	return department.List()
}
