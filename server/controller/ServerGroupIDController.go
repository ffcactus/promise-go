package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// ServerGroupIDController Task controller
type ServerGroupIDController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *ServerGroupIDController) GetResourceName() string {
	return "servergroup"
}

// NewResponse creates a new response DTO.
func (c *ServerGroupIDController) NewResponse() base.ResponseInterface {
	response := new(dto.GetServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *ServerGroupIDController) GetService() base.ServiceInterface {
	return serverGroupService
}
