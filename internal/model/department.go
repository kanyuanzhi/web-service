package model

import "github.com/kanyuanzhi/web-service/global"

type Department struct {
	*DefaultFields
	Name         string `json:"name,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

func (d *Department) Create() (*Department, error) {
	err := global.DB.Create(d).Error
	return d, err
}

func (d *Department) List() ([]*Department, error) {
	var departments []*Department
	err := global.DB.Find(&departments).Error
	return departments, err
}

func (d *Department) Update() (*Department, error) {
	err := global.DB.Model(d).Updates(d).Error
	return d, err
}

func (d *Department) Delete() error {
	err := global.DB.Delete(d).Error
	return err
}
