package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerGroupRootController is ther servergroup controller.
type ServerGroupRootController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerGroupRootController) ResourceName() string {
	return "servergroup"
}

// Request creates a new request DTO.
func (c *ServerGroupRootController) Request() base.PostRequestInterface {
	return new(dto.PostServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerGroupRootController) Response() base.GetResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// Service returns the service.
func (c *ServerGroupRootController) Service() base.CRUDServiceInterface {
	return serverGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerGroupRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
