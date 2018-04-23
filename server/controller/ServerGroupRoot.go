package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerGroupRoot is ther servergroup controller.
type ServerGroupRoot struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerGroupRoot) ResourceName() string {
	return "servergroup"
}

// Request creates a new request DTO.
func (c *ServerGroupRoot) Request() base.PostRequestInterface {
	return new(dto.PostServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerGroupRoot) Response() base.GetResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// Service returns the service.
func (c *ServerGroupRoot) Service() base.CRUDServiceInterface {
	return serverGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerGroupRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
