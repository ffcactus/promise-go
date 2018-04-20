package dto

import (
	"promise/base"
	"promise/server/object/model"
)

// PostServerGroupRequest is the DTO for post server group.
type PostServerGroupRequest struct {
	base.Request
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// IsValid return if the request is valid.
func (dto *PostServerGroupRequest) IsValid() *base.Message {
	return nil
}

// GetDebugName return the name for debug.
func (dto *PostServerGroupRequest) GetDebugName() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto *PostServerGroupRequest) ToModel() base.ModelInterface {
	ret := model.ServerGroup{}
	ret.Category = base.CategoryServerGroup
	ret.Name = dto.Name
	ret.Description = dto.Description
	return &ret
}
