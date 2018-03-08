package dto

import (
	"promise/server/object/model"
)

// PostServerGroupRequest is the DTO for post server group.
type PostServerGroupRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// ToModel will return a model based on DTO.
func (dto *PostServerGroupRequest) ToModel() *model.ServerGroup {
	ret := new(model.ServerGroup)
	ret.Name = dto.Name
	ret.Description = dto.Description
	return ret
}
