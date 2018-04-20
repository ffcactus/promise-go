package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerServerGroupCollectionMember is the DTO used in collection response.
type ServerServerGroupCollectionMember struct {
	base.CollectionMemberResponse
	ServerID       string `json:"ServerID"`
	ServerURI      string `json:"ServerURI"`
	ServerGroupID  string `json:"ServerGroupID"`
	ServerGroupURI string `json:"ServerGroupURI"`
}

// Load will load info from model.
func (dto *ServerServerGroupCollectionMember) Load(i interface{}) error {
	m, ok := i.(*model.ServerServerGroupCollectionMember)
	if !ok {
		log.Error("ServerServerGroupCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.ServerID = m.ServerID
	dto.ServerGroupID = m.ServerGroupID
	dto.ServerURI = base.ToServerURI(m.ServerID)
	dto.ServerGroupURI = base.ToServerGroupURI(m.ServerGroupID)
	return nil
}

// GetServerServerGroupCollectionResponse is the DTO.
type GetServerServerGroupCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetServerServerGroupCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]interface{}, 0)
	for _, v := range m.Members {
		member := ServerServerGroupCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, member)
	}
	return nil
}
