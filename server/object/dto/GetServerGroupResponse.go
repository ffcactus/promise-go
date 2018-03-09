package dto

import (
	"promise/server/object/model"
)

// GetServerGroupResponse is the DTO.
type GetServerGroupResponse struct {
	ResourceResponse
	ID          string `json:"ID"`
	Name        string `json:"URI"`
	Description string `json:"Description"`
}

// Load the data from model.
func (dto *GetServerGroupResponse) Load(m *model.ServerGroup) {
	dto.ID = m.ID
	dto.Name = m.Name
	dto.Description = m.Description
}
