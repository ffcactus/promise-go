package dto

import (
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// ServerMember is DTO.
type ServerMember struct {
	URI    string `json:"URI"`
	Name   string `json:"Name"`
	State  string `json:"State"`
	Health string `json:"Health"`
}

// GetServerCollectionResponse is DTO.
type GetServerCollectionResponse struct {
	Start       int            `json:"Start"`
	Count       int            `json:"Count"`
	Total       int            `json:"Total"`
	Members     []ServerMember `json:"Members"`
	NextPageURI *string        `json:"NextPageURI,omitempty"`
	PrevPageURI *string        `json:"PrevPageURI,omitempty"`
}

// Load will load data from model.
func (dto *GetServerCollectionResponse) Load(m *model.ServerCollection) {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]ServerMember, 0)
	for i := range m.Members {
		dto.Members = append(dto.Members, ServerMember{
			URI:    constvalue.ToServerURI(m.Members[i].ID),
			Name:   m.Members[i].Name,
			State:  m.Members[i].State,
			Health: m.Members[i].Health,
		})
	}
	if m.NextPageURI != "" {
		dto.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		dto.PrevPageURI = &m.PrevPageURI
	}
}
