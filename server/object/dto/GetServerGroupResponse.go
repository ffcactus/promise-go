package dto

import (
	"promise/server/object/constvalue"
	"promise/server/object/model"
)

// GetServerGroupResponse is the DTO.
type GetServerGroupResponse struct {
	ID          string `json:"ID"`
	URI         string `json:"URI"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// Load the data from model.
func (dto *GetServerGroupResponse) Load(m *model.ServerGroup) {
	dto.ID = m.ID
	dto.URI = constvalue.ToServerGroupURI(m.ID)
	dto.Name = m.Name
	dto.Description = m.Description
}
