package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerServerGroupRootController is ther servergroup controller.
type ServerServerGroupRootController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerServerGroupRootController) ResourceName() string {
	return "server-servergroup"
}

// Request creates a new request DTO.
func (c *ServerServerGroupRootController) Request() base.PostRequestInterface {
	return new(dto.PostServerServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerServerGroupRootController) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// Service returns the service.
func (c *ServerServerGroupRootController) Service() base.CRUDServiceInterface {
	return serverServerGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerServerGroupRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
