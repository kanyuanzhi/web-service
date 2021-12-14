package model

import "github.com/kanyuanzhi/web-service/global"

type Department struct {
	*DefaultFields
	Name         string `json:"name,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

func (d *Department) Create() *Department {
	err := global.DB.Create(d).Error
	if err != nil {
		global.Log.Error(err)
	}
	return d
}

func (d *Department) List() []*Department {
	var departments []*Department
	err := global.DB.Find(&departments).Error
	if err != nil {
		global.Log.Error(err)
	}
	return departments
}
