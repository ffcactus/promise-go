package dto

import (
	"promise/server/object/model"
)

// PostGroupRequest is the DTO for post server group.
type PostGroupRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// ToModel will return a model based on DTO.
func (dto *PostGroupRequest) ToModel() *model.Group {
	ret := new(model.Group)
	ret.Name = dto.Name
	ret.Description = dto.Description
	return ret
}
