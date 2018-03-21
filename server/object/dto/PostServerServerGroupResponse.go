package dto

import (
	"promise/common/object/constValue"
	"promise/server/object/model"
)

// PostServerServerGroupResponse is the DTO.
type PostServerServerGroupResponse struct {
	ResourceResponse
	ID             string `json:"ID"`
	URI            string `json:"URI"`
	ServerID       string `json:"ServerID"`
	ServerURI      string `json:"ServerURI"`
	ServerGroupID  string `json:"ServerGroupID"`
	ServerGroupURI string `json:"ServerGroupURI"`
}

// Load the data from model.
func (dto *PostServerServerGroupResponse) Load(m *model.ServerServerGroup) {
	dto.ID = m.ID
	dto.URI = constValue.ToServerServerGroupURI(m.ID)
	dto.ServerID = m.ServerID
	dto.ServerURI = constValue.ToServerURI(m.ServerID)
	dto.ServerGroupID = m.ServerGroupID
	dto.ServerGroupURI = constValue.ToServerGroupURI(m.ServerGroupID)
}
