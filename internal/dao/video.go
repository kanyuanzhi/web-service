package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateVideo(name string, ip string, location string) (*model.Video, error) {
	video := &model.Video{
		Name:     name,
		IP:       ip,
		Location: location,
	}
	return video.Create()
}

func (dao *Dao) ListVideos() ([]*model.Video, error) {
	video := &model.Video{}
	return video.List()
}

func (dao *Dao) UpdateVideo(id uint, name string, ip string, location string) (*model.Video, error) {
	video := &model.Video{
		DefaultFields: &model.DefaultFields{ID: id},
		Name:          name,
		IP:            ip,
		Location:      location,
	}
	return video.Update()
}

func (dao *Dao) DeleteVideo(id uint) error {
	video := &model.Video{
		DefaultFields: &model.DefaultFields{ID: id},
	}
	return video.Delete()
}
