package controller

import (
	"promise/base"
	"promise/pool/object/dto"
)

// IPv4ID Task ID controller
type IPv4ID struct {
}

// ResourceName returns the name this controller handle of.
func (c *IPv4ID) ResourceName() string {
	return "ipv4"
}

// Response creates a new response DTO.
func (c *IPv4ID) Response() base.GetResponseInterface {
	return new(dto.GetIPv4PoolResponse)
}

// Service returns the service.
func (c *IPv4ID) Service() base.CRUDServiceInterface {
	return ipv4Service
}
