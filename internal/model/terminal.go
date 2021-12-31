package model

import (
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/model/terminal_model"
)

type Terminal struct {
	*DefaultFields
	*TerminalManual               `gorm:"embedded"`
	*terminal_model.TerminalBasic `gorm:"embedded"`
}

type TerminalManual struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

func (t *Terminal) Create() (*Terminal, error) {
	err := global.DB.Create(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Terminal) List() ([]*Terminal, error) {
	var terminals []*Terminal
	err := global.DB.Find(&terminals).Error
	if err != nil {
		return nil, err
	}
	return terminals, nil
}

func (t *Terminal) Update() (*Terminal, error) {
	err := global.DB.Model(t).Updates(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Terminal) Delete() error {
	err := global.DB.Delete(t).Error
	if err != nil {
		return err
	}
	return nil
}
