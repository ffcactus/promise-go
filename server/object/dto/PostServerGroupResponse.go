package dto

import (
	"promise/server/object/model"
)

// PostServerGroupResponse is the DTO.
type PostServerGroupResponse struct {
	ResourceResponse
	ID           string `json:"ID"`
	Name          string `json:"URI"`
	Description string `json:"Description"`
}

// Load the data from model.
func (this *PostServerGroupResponse) Load(m *model.ServerGroup) {
	this.ID = m.ID
	this.Name = m.Name
	this.Description = m.Description
}
