package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// ServerServerGroupID Task controller
type ServerServerGroupID struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerServerGroupID) ResourceName() string {
	return "server-servergroup"
}

// Request creates a new request DTO.
func (c *ServerServerGroupID) Request() base.PostRequestInterface {
	return new(dto.PostServerServerGroupRequest)
}

// Response creates a new response DTO.
func (c *ServerServerGroupID) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// Service returns the service.
func (c *ServerServerGroupID) Service() base.CRUDServiceInterface {
	return serverServerGroupService
}
