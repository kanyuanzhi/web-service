package model

import "github.com/kanyuanzhi/web-service/internal/model"

type Terminal struct {
	*model.DefaultFields
	*TerminalManual `gorm:"embedded"`
	*TerminalBasic  `gorm:"embedded"`
}

type TerminalManual struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}
