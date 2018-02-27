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

func (this *PostServerResponse) Load(m *model.Server) {
	this.ID = m.ID
	this.URI = m.URI
	this.PhysicalUUID = m.PhysicalUUID
	this.Address = m.Address
	this.Type = m.Type
}
