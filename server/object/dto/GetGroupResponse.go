package dto

import (
	"promise/server/object/model"
)

// GetGroupResponse is the DTO.
type GetGroupResponse struct {
	ID          string `json:"ID"`
	URI         string `json:"URI"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// Load the data from model.
func (dto *GetGroupResponse) Load(m *model.Group) {
	dto.ID = m.ID
	dto.URI = m.URI
	dto.Name = m.Name
	dto.Description = m.Description
}
