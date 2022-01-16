package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateTerminalGroup(name string, manager string, introduction string) (*model.TerminalGroup, error) {
	department := &model.TerminalGroup{
		Name:         name,
		Manager:      manager,
		Introduction: introduction,
	}
	return department.Create()
}

func (dao *Dao) ListTerminalGroups() ([]*model.TerminalGroup, error) {
	department := &model.TerminalGroup{}
	return department.List()
}

func (dao *Dao) UpdateTerminalGroup(id uint, name string, manager string, introduction string) (*model.TerminalGroup, error) {
	department := &model.TerminalGroup{
		DefaultFields: &model.DefaultFields{ID: id},
		Name:          name,
		Manager:       manager,
		Introduction:  introduction,
	}
	return department.Update()
}

func (dao *Dao) DeleteTerminalGroup(id uint) error {
	department := &model.TerminalGroup{
		DefaultFields: &model.DefaultFields{ID: id},
	}
	return department.Delete()
}
