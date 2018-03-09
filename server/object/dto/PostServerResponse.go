package dto

import (
	"promise/server/object/model"
)

type PostServerResponse struct {
	ResourceResponse
	ID           string `json:"ID"`
	URI          string `json:"URI"`
	PhysicalUUID string `json:"PhysicalUUID"`
	Address      string `json:"Address"`
	Type         string `json:"Type"`
}

func (dto *PostServerResponse) Load(m *model.Server) {
	dto.ID = m.ID
	dto.URI = m.URI
	dto.PhysicalUUID = m.PhysicalUUID
	dto.Address = m.Address
	dto.Type = m.Type
}
