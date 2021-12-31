package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
	"github.com/kanyuanzhi/web-service/internal/model/terminal_model"
)

func (dao *Dao) CreateTerminal(name string, manager string, mac string) (*model.Terminal, error) {
	terminal := &model.Terminal{
		TerminalManual: &model.TerminalManual{
			Name:    name,
			Manager: manager,
		},
		TerminalBasic: &terminal_model.TerminalBasic{
			NetBasic: &terminal_model.TerminalNetBasic{
				Mac: mac,
			},
		},
	}
	return terminal.Create()
}

func (dao *Dao) ListTerminals() ([]*model.Terminal, error) {
	terminal := &model.Terminal{}
	return terminal.List()
}

func (dao *Dao) UpdateTerminal(id uint, name string, manager string, mac string) (*model.Terminal, error) {
	terminal := &model.Terminal{
		DefaultFields: &model.DefaultFields{
			ID: id,
		},
		TerminalManual: &model.TerminalManual{
			Name:    name,
			Manager: manager,
		},
		TerminalBasic: &terminal_model.TerminalBasic{
			NetBasic: &terminal_model.TerminalNetBasic{
				Mac: mac,
			},
		},
	}
	return terminal.Update()
}

func (dao *Dao) DeleteTerminal(id uint) error {
	terminal := &model.Terminal{
		DefaultFields: &model.DefaultFields{
			ID: id,
		},
	}
	return terminal.Delete()
}
