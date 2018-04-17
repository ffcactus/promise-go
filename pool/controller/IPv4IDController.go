package controller

import (
	"promise/base"
	"promise/pool/object/dto"
)

// IPvID4Controller Task controller
type IPvID4Controller struct {
}

// GetResourceName returns the name this controller handle of.
func (c *IPvID4Controller) GetResourceName() string {
	return "ipv4"
}

// NewResponse creates a new response DTO.
func (c *IPvID4Controller) NewResponse() base.ResponseInterface {
	response := new(dto.GetIPv4PoolResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *IPvID4Controller) GetService() base.ServiceInterface {
	return ipv4Service
}
