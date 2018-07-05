package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/director/object/model"
)

// NodeCollectionMember is the DTO used in collection response.
type NodeCollectionMember struct {
	base.CollectionMemberResponse
	Hostname string `json:"Hostname"`
	Status string `json:"Status"`
	Availibility string `json:"Availibility"`
	ManagerStatus string `json:"ManagerStatus"`
}

// Load will load info from model.
func (dto *NodeCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.NodeCollectionMember)
	if !ok {
		log.Error("NodeCollectionMember.Load() failed, convert data failed.")
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Hostname = m.Hostname
	dto.Status = m.Status
	dto.Availibility = m.Availibility
	dto.ManagerStatus = m.ManagerStatus
	return nil
}

// GetNodeCollectionResponse is the DTO.
type GetNodeCollectionResponse struct {
	base.CollectionResponse
}

// Load will load from model.
func (dto *GetNodeCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = make([]base.CollectionMemberResponseInterface, 0)
	for _, v := range m.Members {
		member := NodeCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}
