package dto

import (
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// ServerGroupMember is the Members property in DTO.
type ServerGroupMember struct {
	URI  string `json:"URI"`
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// GetServerGroupCollectionResponse is the DTO.
type GetServerGroupCollectionResponse struct {
	Start       int                 `json:"Start"`
	Count       int                 `json:"Count"`
	Total       int                 `json:"Total"`
	Members     []ServerGroupMember `json:"Members"`
	NextPageURI *string             `json:"NextPageURI,omitempty"`
	PrevPageURI *string             `json:"PrevPageURI,omitempty"`
}

// Load will load from model.
func (dto *GetServerGroupCollectionResponse) Load(m *model.ServerGroupCollection) {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]ServerGroupMember, 0)
	for i := range m.Members {
		dto.Members = append(dto.Members, ServerGroupMember{
			URI:  constvalue.ToServerGroupURI(m.Members[i].ID),
			ID:   m.Members[i].ID,
			Name: m.Members[i].Name,
		})
	}
	if m.NextPageURI != "" {
		dto.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		dto.PrevPageURI = &m.PrevPageURI
	}
}
