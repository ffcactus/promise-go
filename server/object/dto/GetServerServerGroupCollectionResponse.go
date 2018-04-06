package dto

import (
	"promise/common/object/constvalue"
	"promise/server/object/model"
)

// ServerServerGroupMember is the Members property in DTO.
type ServerServerGroupMember struct {
	URI            string `json:"URI"`
	ID             string `json:"ID"`
	ServerID       string `json:"ServerID"`
	ServerURI      string `json:"ServerURI"`
	ServerGroupID  string `json:"ServerGroupID"`
	ServerGroupURI string `json:"ServerGroupURI"`
}

// GetServerServerGroupCollectionResponse is the DTO.
type GetServerServerGroupCollectionResponse struct {
	Start       int64                     `json:"Start"`
	Count       int64                     `json:"Count"`
	Total       int64                     `json:"Total"`
	Members     []ServerServerGroupMember `json:"Members"`
	NextPageURI *string                   `json:"NextPageURI,omitempty"`
	PrevPageURI *string                   `json:"PrevPageURI,omitempty"`
}

// Load will load from model.
func (dto *GetServerServerGroupCollectionResponse) Load(m *model.ServerServerGroupCollection) {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]ServerServerGroupMember, 0)
	for i := range m.Members {
		dto.Members = append(dto.Members, ServerServerGroupMember{
			URI:            constvalue.ToServerServerGroupURI(m.Members[i].ID),
			ID:             m.Members[i].ID,
			ServerID:       m.Members[i].ServerID,
			ServerURI:      constvalue.ToServerURI(m.Members[i].ServerID),
			ServerGroupID:  m.Members[i].ServerGroupID,
			ServerGroupURI: constvalue.ToServerGroupURI(m.Members[i].ServerGroupID),
		})
	}
	if m.NextPageURI != "" {
		dto.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		dto.PrevPageURI = &m.PrevPageURI
	}
}
