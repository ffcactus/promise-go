package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// AdapterModelCollectionMember is the DTO used in collection response.
type AdapterModelCollectionMember struct {
	base.CollectionMemberResponse
	Name string `json:"Name"`
	Type string `json:"Type"`
}

// Load will load info from model.
func (dto *AdapterModelCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.AdapterModelCollectionMember)
	if !ok {
		log.Error("AdapterModelCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	dto.Type = m.Type
	return nil
}

// GetAdapterModelCollectionResponse is the DTO.
type GetAdapterModelCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetAdapterModelCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]base.CollectionMemberResponseInterface, 0)
	for _, v := range m.Members {
		member := AdapterModelCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
