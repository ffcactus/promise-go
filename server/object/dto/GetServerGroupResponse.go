package dto

import (
	"promise/common/object/constValue"
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
	dto.URI = constValue.ToServerGroupURI(m.ID)
	dto.Name = m.Name
	dto.Description = m.Description
}
