package dto

import (
	"promise/server/object/model"
)

// GroupMember is the Members property in DTO.
type GroupMember struct {
	URI  string `json:"URI"`
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// GetGroupCollectionResponse is the DTO.
type GetGroupCollectionResponse struct {
	Start       int           `json:"Start"`
	Count       int           `json:"Count"`
	Total       int           `json:"Total"`
	Members     []GroupMember `json:"Members"`
	NextPageURI *string       `json:"NextPageURI,omitempty"`
	PrevPageURI *string       `json:"PrevPageURI,omitempty"`
}

// Load will load from model.
func (dto *GetGroupCollectionResponse) Load(m *model.GroupCollection) {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]GroupMember, 0)
	for i := range m.Members {
		dto.Members = append(dto.Members, GroupMember{
			URI:  m.Members[i].URI,
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
