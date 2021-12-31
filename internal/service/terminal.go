package service

import "github.com/kanyuanzhi/web-service/internal/model"

type CreateTerminalRequest struct {
	Name    string `json:"name" form:"name"`
	Manager string `json:"manager" form:"manager"`
	Mac     string `json:"mac" form:"mac"`
}

func (s *Service) CreateTerminal(param *CreateTerminalRequest) (*model.Terminal, error) {
	return s.dao.CreateTerminal(param.Name, param.Manager, param.Mac)
}

type ListTerminalsRequest struct{}

func (s *Service) ListTerminals() ([]*model.Terminal, error) {
	return s.dao.ListTerminals()
}

type UpdateTerminalRequest struct {
	ID      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Manager string `json:"manager" form:"manager"`
	Mac     string `json:"mac" form:"mac"`
}

func (s *Service) UpdateTerminal(param *UpdateTerminalRequest) (*model.Terminal, error) {
	return s.dao.UpdateTerminal(param.ID, param.Name, param.Manager, param.Mac)
}

type DeleteTerminalRequest struct {
	ID uint `json:"id" form:"id"`
}

func (s *Service) DeleteTerminal(param *DeleteTerminalRequest) error {
	return s.dao.DeleteTerminal(param.ID)
}
