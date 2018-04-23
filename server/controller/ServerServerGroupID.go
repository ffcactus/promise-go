package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

var (
	// ipv4Service is the service used in Student controller.
	serverServerGroupService = &base.CRUDService{
		TemplateImpl: new(service.ServerServerGroup),
	}
)

// ServerServerGroupID Task controller
type ServerServerGroupID struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerServerGroupID) ResourceName() string {
	return "server-servergroup"
}

// NewResponse creates a new response DTO.
func (c *ServerServerGroupID) NewResponse() base.PostRequestInterface {
	response := new(dto.GetServerServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// Response creates a new response DTO.
func (c *ServerServerGroupID) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// Service returns the service.
func (c *ServerServerGroupID) Service() base.CRUDServiceInterface {
	return serverServerGroupService
}
