package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerGroupCollectionMember is the DTO used in collection response.
type ServerGroupCollectionMember struct {
	base.CollectionMemberResponse
	Name string `json:"Name"`
}

// Load will load info from model.
func (dto *ServerGroupCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.ServerGroupCollectionMember)
	if !ok {
		log.Error("ServerGroupCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	return nil
}

// GetServerGroupCollectionResponse is the DTO.
type GetServerGroupCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetServerGroupCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	for _, v := range m.Members {
		member := ServerGroupCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
