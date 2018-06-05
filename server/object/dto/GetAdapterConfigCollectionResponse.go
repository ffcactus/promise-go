package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// AdapterConfigCollectionMember is the DTO used in collection response.
type AdapterConfigCollectionMember struct {
	base.CollectionMemberResponse
	Name string `json:"Name"`
}

// Load will load info from model.
func (dto *AdapterConfigCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.AdapterConfigCollectionMember)
	if !ok {
		log.Error("AdapterConfigCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	return nil
}

// GetAdapterConfigCollectionResponse is the DTO.
type GetAdapterConfigCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetAdapterConfigCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]base.CollectionMemberResponseInterface, 0)
	for _, v := range m.Members {
		member := AdapterConfigCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
