package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

var (
	// ipv4Service is the service used in Student controller.
	serverServerGroupService = &base.CRUDService{
		TemplateImpl: new(service.ServerServerGroup),
	}
)

// ServerServerGroupRoot is ther servergroup controller.
type ServerServerGroupRoot struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerServerGroupRoot) ResourceName() string {
	return "server-servergroup"
}

// Request creates a new request DTO.
func (c *ServerServerGroupRoot) Request() base.PostRequestInterface {
	return new(dto.PostServerServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerServerGroupRoot) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// Service returns the service.
func (c *ServerServerGroupRoot) Service() base.CRUDServiceInterface {
	return serverServerGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerServerGroupRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
