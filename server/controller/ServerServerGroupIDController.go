package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// ServerServerGroupIDController Task controller
type ServerServerGroupIDController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *ServerServerGroupIDController) GetResourceName() string {
	return "server-servergroup"
}

// NewResponse creates a new response DTO.
func (c *ServerServerGroupIDController) NewResponse() base.ResponseInterface {
	response := new(dto.GetServerServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *ServerServerGroupIDController) GetService() base.ServiceInterface {
	return serverServerServerGroupService
}
