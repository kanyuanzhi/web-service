package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateTerminalGroup(name string, manager string, introduction string) (*model.TerminalGroup, error) {
	terminalGroup := &model.TerminalGroup{
		Name:         name,
		Manager:      manager,
		Introduction: introduction,
	}
	return terminalGroup.Create()
}

func (dao *Dao) ListTerminalGroups() ([]*model.TerminalGroup, error) {
	terminalGroup := &model.TerminalGroup{}
	return terminalGroup.List()
}

func (dao *Dao) UpdateTerminalGroup(id uint, name string, manager string, introduction string) (*model.TerminalGroup, error) {
	terminalGroup := &model.TerminalGroup{
		DefaultFields: &model.DefaultFields{ID: id},
		Name:          name,
		Manager:       manager,
		Introduction:  introduction,
	}
	return terminalGroup.Update()
}

func (dao *Dao) DeleteTerminalGroup(id uint) error {
	terminalGroup := &model.TerminalGroup{
		DefaultFields: &model.DefaultFields{ID: id},
	}
	return terminalGroup.Delete()
}
