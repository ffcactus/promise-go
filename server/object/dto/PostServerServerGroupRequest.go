package dto

import (
	"promise/server/object/model"
)

// PostServerServerGroupRequest is the DTO for post server-group.
type PostServerServerGroupRequest struct {
	ServerID      string `json:"ServerID"`
	ServerGroupID string `json:"GroupID"`
}

// ToModel will return a model based on DTO.
func (dto *PostServerServerGroupRequest) ToModel() *model.ServerServerGroup {
	ret := new(model.ServerServerGroup)
	ret.ServerID = dto.ServerID
	ret.ServerGroupID = dto.ServerGroupID
	return ret
}
