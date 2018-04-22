package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

var (
	serverGroupService = &base.CRUDService{
		TemplateImpl: new(service.ServerGroup),
	}
)

// ServerGroupIDController Task controller
type ServerGroupIDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerGroupIDController) ResourceName() string {
	return "servergroup"
}

// Request creates a new request DTO.
func (c *ServerGroupIDController) Request() base.PostRequestInterface {
	return new(dto.PostServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerGroupIDController) Response() base.GetResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// Service returns the service.
func (c *ServerGroupIDController) Service() base.CRUDServiceInterface {
	return serverGroupService
}

// GetService returns the service.
func (c *ServerGroupIDController) GetService() base.ServiceInterface {
	return serverGroupService
}
