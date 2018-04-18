package controller

import (
	"promise/base"
	"promise/pool/object/dto"
)

// IPv4IDController Task controller
type IPv4IDController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *IPv4IDController) GetResourceName() string {
	return "ipv4"
}

// NewResponse creates a new response DTO.
func (c *IPv4IDController) NewResponse() base.ResponseInterface {
	response := new(dto.GetIPv4PoolResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *IPv4IDController) GetService() base.ServiceInterface {
	return ipv4Service
}
