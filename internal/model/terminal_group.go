package model

import "github.com/kanyuanzhi/web-service/global"

type TerminalGroup struct {
	*DefaultFields
	Name         string `json:"name,omitempty"`
	Manager      string `json:"manager,omitempty"`
	Introduction string `json:"introduction,omitempty"`

	Members []uint `json:"members" gorm:"-"`
}

func (tg *TerminalGroup) TableName() string {
	return "terminal_groups"
}

func (tg *TerminalGroup) Create() (*TerminalGroup, error) {
	err := global.DB.Create(tg).Error
	if err != nil {
		return nil, err
	}
	return tg, nil
}

func (tg *TerminalGroup) List() ([]*TerminalGroup, error) {
	var TerminalGroups []*TerminalGroup
	err := global.DB.Find(&TerminalGroups).Error
	if err != nil {
		return nil, err
	}
	return TerminalGroups, nil
}

func (tg *TerminalGroup) Update() (*TerminalGroup, error) {
	err := global.DB.Model(tg).Updates(tg).Error
	if err != nil {
		return nil, err
	}
	return tg, nil
}

func (tg *TerminalGroup) Delete() error {
	err := global.DB.Delete(tg).Error
	return err
}
