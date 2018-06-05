package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// AdapterModelID Task controller
type AdapterModelID struct {
}

// ResourceName returns the name this controller handle of.
func (c *AdapterModelID) ResourceName() string {
	return "servergroup"
}

// Request creates a new request DTO.
func (c *AdapterModelID) Request() base.PostRequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *AdapterModelID) Response() base.GetResponseInterface {
	return new(dto.GetAdapterModelResponse)
}

// Service returns the service.
func (c *AdapterModelID) Service() base.CRUDServiceInterface {
	return serverGroupService
}
