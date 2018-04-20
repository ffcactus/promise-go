package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerCollectionMember is the DTO used in collection response.
type ServerCollectionMember struct {
	base.CollectionMemberResponse
	Name   string `json:"Name"`
	State  string `json:"State"`
	Health string `json:"Health"`
}

// Load will load info from model.
func (dto *ServerCollectionMember) Load(i interface{}) error {
	m, ok := i.(*model.ServerCollectionMember)
	if !ok {
		log.Error("ServerCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	dto.State = m.State
	dto.Health = m.Health
	return nil
}

// GetServerCollectionResponse is DTO.
type GetServerCollectionResponse struct {
	base.CollectionResponse
}

// Load will load data from model.
func (dto *GetServerCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]interface{}, 0)
	for _, v := range m.Members {
		member := ServerCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, member)
	}
	return nil
}
