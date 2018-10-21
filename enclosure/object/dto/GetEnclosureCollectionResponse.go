package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// EnclosureCollectionMember is the DTO used in collection response.
type EnclosureCollectionMember struct {
	base.CollectionMemberResponse
	Name        string
	Description string
	State       string
	Health      string
}

// Load will load info from model.
func (dto *EnclosureCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.EnclosureCollectionMember)
	if !ok {
		log.Error("EnclosureCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	return nil
}

// GetEnclosureCollectionResponse is the DTO.
type GetEnclosureCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetEnclosureCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]base.CollectionMemberResponseInterface, 0)
	for _, v := range m.Members {
		member := EnclosureCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
