package dto

import (
	commonModel "promise/common/object/model"
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
}

// Load will load from model.
func (dto *GetIPv4PoolCollectionResponse) Load(m commonModel.PromiseCollectionInterface) {
	dto.Start = m.GetStart()
	dto.Count = m.GetCount()
	dto.Total = m.GetTotal()
	dto.Members = make([]IPv4PoolMember, 0)
	members := m.GetMembers()
	for i := range members {
		dto.Members = append(dto.Members, IPv4PoolMember{
			URI:  constvalue.ToIDPoolIPv4URI(members[i].ID),
			ID:   members[i].ID,
			Name: members[i].Name,
		})
	}
}
