package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type CreateVideoRequest struct {
	Name     string `json:"name" form:"name" `
	IP       string `json:"ip" form:"ip"`
	Location string `json:"location" form:"location"`
}

func (s *Service) CreateVideo(param *CreateVideoRequest) (*model.Video, error) {
	return s.dao.CreateVideo(param.Name, param.IP, param.Location)
}

type ListVideosRequest struct{}

func (s *Service) ListVideos() ([]*model.Video, error) {
	groups, err := s.dao.ListVideos()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

type UpdateVideoRequest struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" `
	IP       string `json:"ip" form:"ip"`
	Location string `json:"location" form:"location"`
}

func (s *Service) UpdateVideo(param *UpdateVideoRequest) (*model.Video, error) {
	return s.dao.UpdateVideo(param.ID, param.Name, param.IP, param.Location)
}

type DeleteVideoRequest struct {
	ID uint `json:"id" form:"id"`
}

func (s *Service) DeleteVideo(param *DeleteVideoRequest) error {
	err := s.dao.DeleteVideo(param.ID)
	if err != nil {
		return err
	}
	return nil
}
