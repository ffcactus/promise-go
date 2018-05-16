package dto

import (
	"promise/base"
	"promise/server/object/model"
)

// PostServerServerGroupRequest is the DTO for post server-group.
type PostServerServerGroupRequest struct {
	ServerID      string `json:"ServerID"`
	ServerGroupID string `json:"ServerGroupID"`
}

// NewInstance creates a new instance.
func (dto *PostServerServerGroupRequest) NewInstance() base.RequestInterface {
	return new(PostServerServerGroupRequest)
}

// IsValid return if the request is valid.
func (dto *PostServerServerGroupRequest) IsValid() *base.Message {
	if dto.ServerID == "" || dto.ServerGroupID == "" {
		return base.NewMessageInvalidRequest()
	} 
	return nil
}

// DebugInfo return the name for debug.
func (dto *PostServerServerGroupRequest) DebugInfo() string {
	return dto.ServerID + " " + dto.ServerGroupID
}

// ToModel convert the DTO to model.
func (dto *PostServerServerGroupRequest) ToModel() base.ModelInterface {
	ret := model.ServerServerGroup{}
	ret.Category = base.CategoryServerServerGroup
	ret.ServerID = dto.ServerID
	ret.ServerGroupID = dto.ServerGroupID
	return &ret
}
