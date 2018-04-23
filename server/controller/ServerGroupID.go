package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// ServerGroupID Task controller
type ServerGroupID struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerGroupID) ResourceName() string {
	return "servergroup"
}

// Request creates a new request DTO.
func (c *ServerGroupID) Request() base.PostRequestInterface {
	return new(dto.PostServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerGroupID) Response() base.GetResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// Service returns the service.
func (c *ServerGroupID) Service() base.CRUDServiceInterface {
	return serverGroupService
}
