package dao

import (
	"github.com/kanyuanzhi/web-service/internal/model"
)

func (dao *Dao) CreateTerminalGroupAssociations(groupID uint, terminalIDs []uint) ([]*model.TerminalGroupAssociation, error) {
	association := &model.TerminalGroupAssociation{}
	var associations []*model.TerminalGroupAssociation
	for _, terminalID := range terminalIDs {
		associations = append(associations,
			&model.TerminalGroupAssociation{GroupID: groupID, TerminalID: terminalID})
	}
	return association.CreateMany(associations)
}

func (dao *Dao) GetTerminalGroupAssociationsByGroupID(groupID uint) ([]*model.TerminalGroupAssociation, error) {
	association := &model.TerminalGroupAssociation{GroupID: groupID}
	return association.Get()
}

func (dao *Dao) DeleteTerminalGroupAssociationsByGroupID(groupID uint) error {
	association := &model.TerminalGroupAssociation{GroupID: groupID}
	return association.Delete()
}

func (dao *Dao) DeleteTerminalGroupAssociationsByTerminalID(terminalID uint) error {
	association := &model.TerminalGroupAssociation{TerminalID: terminalID}
	return association.Delete()
}
