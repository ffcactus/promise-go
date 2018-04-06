package dto

import (
	"promise/common/object/constvalue"
	"promise/pool/object/model"
)

// IPv4PoolMember is the Members property in DTO.
type IPv4PoolMember struct {
	URI  string `json:"URI"`
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// GetIPv4PoolCollectionResponse is the DTO.
type GetIPv4PoolCollectionResponse struct {
	Start       int64            `json:"Start"`
	Count       int64            `json:"Count"`
	Total       int64            `json:"Total"`
	Members     []IPv4PoolMember `json:"Members"`
	NextPageURI *string          `json:"NextPageURI,omitempty"`
	PrevPageURI *string          `json:"PrevPageURI,omitempty"`
}

// Load will load from model.
func (dto *GetIPv4PoolCollectionResponse) Load(m *model.IPv4PoolCollection) {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]IPv4PoolMember, 0)
	for i := range m.Members {
		dto.Members = append(dto.Members, IPv4PoolMember{
			URI:  constvalue.ToIDPoolIPv4URI(m.Members[i].ID),
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
