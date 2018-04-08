package dto

import (
	log "github.com/sirupsen/logrus"
	commonConstError "promise/common/object/consterror"
	"promise/common/object/constvalue"
	commonModel "promise/common/object/model"
	"promise/pool/object/model"
)

// IPv4PoolMember is the Members property in DTO.
type IPv4PoolMember struct {
	URI  string `json:"URI"`
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// Load will load from model.
func (dto *IPv4PoolMember) Load(i commonModel.PromiseMemberInterface) error {
	m, ok := i.(*model.IPv4PoolMember)
	if !ok {
		log.WithFields(log.Fields{
			"from": "PromiseMemberInterface",
			"to":   "IPv4PoolMember",
		}).Fatal("Data convert error!")
		return commonConstError.ErrorDataConvert
	}
	dto.ID = m.ID
	dto.URI = constvalue.ToIDPoolIPv4URI(m.ID)
	dto.Name = m.Name
	return nil
}

// GetIPv4PoolCollectionResponse is the DTO.
type GetIPv4PoolCollectionResponse struct {
	Start   int64            `json:"Start"`
	Count   int64            `json:"Count"`
	Total   int64            `json:"Total"`
	Members []IPv4PoolMember `json:"Members"`
}

// Load will load from model.
func (dto *GetIPv4PoolCollectionResponse) Load(m commonModel.PromiseCollectionInterface) {
	dto.Start = m.GetStart()
	dto.Count = m.GetCount()
	dto.Total = m.GetTotal()
	dto.Members = make([]IPv4PoolMember, 0)
	for _, v := range m.GetMembers() {
		member := IPv4PoolMember{}
		member.Load(v)
		dto.Members = append(dto.Members, member)
	}
}
