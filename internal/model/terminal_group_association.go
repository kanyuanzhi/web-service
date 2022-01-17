package model

import (
	"github.com/kanyuanzhi/web-service/global"
)

type TerminalGroupAssociation struct {
	*DefaultFields
	GroupID    uint `json:"group_id"`
	TerminalID uint `json:"terminal_id"`
}

func (tga *TerminalGroupAssociation) TableName() string {
	return "terminal_group_associations"
}

func (tga *TerminalGroupAssociation) CreateMany(associations []*TerminalGroupAssociation) ([]*TerminalGroupAssociation, error) {
	err := global.DB.Create(&associations).Error
	return associations, err
}

func (tga *TerminalGroupAssociation) Get() ([]*TerminalGroupAssociation, error) {
	var associations []*TerminalGroupAssociation
	err := global.DB.Where(tga).Find(&associations).Error
	return associations, err
}

func (tga *TerminalGroupAssociation) Delete() error {
	err := global.DB.Where(tga).Delete(tga).Error
	return err
}
