package model

import "github.com/kanyuanzhi/web-service/global"

type Video struct {
	*DefaultFields
	Name     string `json:"name,omitempty"`
	IP       string `json:"ip,omitempty"`
	Location string `json:"location,omitempty"`
}

func (v *Video) TableName() string {
	return "videos"
}

func (v *Video) Create() (*Video, error) {
	err := global.DB.Create(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *Video) List() ([]*Video, error) {
	var Videos []*Video
	err := global.DB.Find(&Videos).Error
	if err != nil {
		return nil, err
	}
	return Videos, nil
}

func (v *Video) Update() (*Video, error) {
	err := global.DB.Model(v).Updates(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *Video) Delete() error {
	err := global.DB.Delete(v).Error
	return err
}
