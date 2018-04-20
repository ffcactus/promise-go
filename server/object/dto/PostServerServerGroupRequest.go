package dto

import (
	"promise/base"
	"promise/server/object/model"
)

// PostServerServerGroupRequest is the DTO for post server-group.
type PostServerServerGroupRequest struct {
	base.Request
	ServerID      string `json:"ServerID"`
	ServerGroupID string `json:"ServerGroupID"`
}

// IsValid return if the request is valid.
func (dto *PostServerServerGroupRequest) IsValid() *base.Message {
	return nil
}

// ToModel convert the DTO to model.
func (dto *PostServerServerGroupRequest) ToModel() base.ModelInterface {
	ret := model.ServerServerGroup{}
	ret.ServerID = dto.ServerID
	ret.ServerGroupID = dto.ServerGroupID
	return &ret
}
