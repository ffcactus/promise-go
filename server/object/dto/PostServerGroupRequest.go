package dto

import (
	"promise/base"
	"promise/server/object/model"
)

// PostServerGroupRequest is the DTO for post server group.
type PostServerGroupRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// NewInstance creates a new instance.
func (dto *PostServerGroupRequest) NewInstance() base.RequestInterface {
	return new(PostServerGroupRequest)
}

// IsValid return if the request is valid.
func (dto *PostServerGroupRequest) IsValid() *base.ErrorResponse {
	return nil
}

// String return the name for debug.
func (dto PostServerGroupRequest) String() string {
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
