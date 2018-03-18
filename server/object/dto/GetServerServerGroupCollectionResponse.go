package dto

import (
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// ServerServerGroupMember is the Members property in DTO.
type ServerServerGroupMember struct {
	URI string `json:"URI"`
	ID  string `json:"ID"`
}

// GetServerServerGroupCollectionResponse is the DTO.
type GetServerServerGroupCollectionResponse struct {
	Start       int                       `json:"Start"`
	Count       int                       `json:"Count"`
	Total       int                       `json:"Total"`
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
			URI: constvalue.ToServerServerGroupURI(m.Members[i].ID),
			ID:  m.Members[i].ID,
		})
	}
	if m.NextPageURI != "" {
		dto.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		dto.PrevPageURI = &m.PrevPageURI
	}
}
