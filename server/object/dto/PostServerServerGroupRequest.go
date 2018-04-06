package dto

import (
	commonDTO "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/server/object/model"
)

// PostServerServerGroupRequest is the DTO for post server-group.
type PostServerServerGroupRequest struct {
	commonDTO.PromiseRequest
	ServerID      string `json:"ServerID"`
	ServerGroupID string `json:"ServerGroupID"`
}

// ToModel will return a model based on DTO.
func (dto *PostServerServerGroupRequest) ToModel() *model.ServerServerGroup {
	ret := new(model.ServerServerGroup)
	ret.ServerID = dto.ServerID
	ret.ServerGroupID = dto.ServerGroupID
	return ret
}

// Validate the request.
func (dto *PostServerServerGroupRequest) Validate() *commonMessage.Message {
	return nil
}
