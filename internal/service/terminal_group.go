package service

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

type CreateTerminalGroupRequest struct {
	Name         string `json:"name" form:"name" `
	Manager      string `json:"manager" form:"manager"`
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) CreateTerminalGroup(param *CreateTerminalGroupRequest) (*model.TerminalGroup, error) {
	return s.dao.CreateTerminalGroup(param.Name, param.Manager, param.Introduction)
}

type ListTerminalGroupsRequest struct{}

func (s *Service) ListTerminalGroups() ([]*model.TerminalGroup, error) {
	groups, err := s.dao.ListTerminalGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		terminalGroupAssociations, _ := s.dao.GetTerminalGroupAssociationsByGroupID(group.ID)
		var members []uint
		for _, association := range terminalGroupAssociations {
			members = append(members, association.TerminalID)
		}
		if members == nil {
			group.Members = []uint{}
		} else {
			group.Members = members
		}
	}
	return groups, nil
}

type UpdateTerminalGroupRequest struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name" `
	Manager      string `json:"manager" form:"manager"`
	Introduction string `json:"introduction" form:"introduction"`
}

func (s *Service) UpdateTerminalGroup(param *UpdateTerminalGroupRequest) (*model.TerminalGroup, error) {
	return s.dao.UpdateTerminalGroup(param.ID, param.Name, param.Manager, param.Introduction)
}

type DeleteTerminalGroupRequest struct {
	ID uint `json:"id" form:"id"`
}

func (s *Service) DeleteTerminalGroup(param *DeleteTerminalGroupRequest) error {
	err := s.dao.DeleteTerminalGroup(param.ID)
	if err != nil {
		return err
	}
	err = s.dao.DeleteTerminalGroupAssociationsByGroupID(param.ID)
	if err != nil {
		return err
	}
	return nil
}

type UpdateTerminalGroupAssociationsRequest struct {
	GroupID     uint   `json:"group_id" form:"group_id"`
	TerminalIDs []uint `json:"terminal_ids" form:"terminal_ids"`
}

func (s *Service) UpdateTerminalGroupAssociations(param *UpdateTerminalGroupAssociationsRequest) error {
	err := s.dao.DeleteTerminalGroupAssociationsByGroupID(param.GroupID)
	if err != nil {
		return err
	}
	if len(param.TerminalIDs) != 0 {
		_, err = s.dao.CreateTerminalGroupAssociations(param.GroupID, param.TerminalIDs)
		if err != nil {
			return err
		}
	}
	return nil
}
