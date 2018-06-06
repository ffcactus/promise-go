package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// AdapterConfigID Task controller
type AdapterConfigID struct {
}

// ResourceName returns the name this controller handle of.
func (c *AdapterConfigID) ResourceName() string {
	return "adapterconfig"
}

// Request creates a new request DTO.
func (c *AdapterConfigID) Request() base.PostRequestInterface {
	return new(dto.PostAdapterConfigRequest)
}

// Response creates a new response DTO.
func (c *AdapterConfigID) Response() base.GetResponseInterface {
	return new(dto.GetAdapterConfigResponse)
}

// Service returns the service.
func (c *AdapterConfigID) Service() base.CRUDServiceInterface {
	return adapterConfigService
}
