package dto

import (
	"promise/common/category"
	commonDTO "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/server/object/model"
)

// PostServerGroupRequest is the DTO for post server group.
type PostServerGroupRequest struct {
	commonDTO.PromiseRequest
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// Validate the request.
func (dto *PostServerGroupRequest) Validate() *commonMessage.Message {
	return nil
}

// ToModel will return a model based on DTO.
func (dto *PostServerGroupRequest) ToModel() *model.ServerGroup {
	ret := new(model.ServerGroup)
	ret.Category = category.ServerGroup
	ret.Name = dto.Name
	ret.Description = dto.Description
	return ret
}
