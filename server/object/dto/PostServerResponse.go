package dto

import (
	"promise/common/object/constValue"
	"promise/server/object/model"
)

// PostServerResponse is DTO.
type PostServerResponse struct {
	ResourceResponse
	ID           string `json:"ID"`
	URI          string `json:"URI"`
	PhysicalUUID string `json:"PhysicalUUID"`
	Hostname     string `json:"Hostname"`
	Type         string `json:"Type"`
}

// Load will load data from model.
func (dto *PostServerResponse) Load(m *model.Server) {
	dto.ID = m.ID
	dto.URI = constValue.ToServerURI(m.ID)
	dto.PhysicalUUID = m.PhysicalUUID
	dto.Hostname = m.Hostname
	dto.Type = m.Type
}
