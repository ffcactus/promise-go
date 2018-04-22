package controller

import (
	"promise/base"
	"promise/pool/object/dto"
)

// IPv4IDController Task controller
type IPv4IDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *IPv4IDController) ResourceName() string {
	return "ipv4"
}

// Response creates a new response DTO.
func (c *IPv4IDController) Response() base.GetResponseInterface {
	return new(dto.GetIPv4PoolResponse)
}

// Service returns the service.
func (c *IPv4IDController) Service() base.CRUDServiceInterface {
	return ipv4Service
}
