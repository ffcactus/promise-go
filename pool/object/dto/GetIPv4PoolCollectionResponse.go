package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/model"
)

// IPv4PoolCollectionMember is the DTO used in collection response.
type IPv4PoolCollectionMember struct {
	base.CollectionMemberResponse
	Name string `json:"Name"`
}

// Load will load info from model.
func (dto *IPv4PoolCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.IPv4PoolCollectionMember)
	if !ok {
		log.Error("IPv4PoolCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	return nil
}

// GetIPv4PoolCollectionResponse is the DTO.
type GetIPv4PoolCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetIPv4PoolCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]base.CollectionMemberResponseInterface, 0)
	for _, v := range m.Members {
		member := IPv4PoolCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
