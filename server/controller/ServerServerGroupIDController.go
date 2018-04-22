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

// ServerServerGroupIDController Task controller
type ServerServerGroupIDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerServerGroupIDController) ResourceName() string {
	return "server-servergroup"
}

// NewResponse creates a new response DTO.
func (c *ServerServerGroupIDController) NewResponse() base.PostRequestInterface {
	response := new(dto.GetServerServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// Response creates a new response DTO.
func (c *ServerServerGroupIDController) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// Service returns the service.
func (c *ServerServerGroupIDController) Service() base.CRUDServiceInterface {
	return serverServerGroupService
}
