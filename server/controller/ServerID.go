package controller

import (
	"promise/base"
	"promise/server/object/dto"
)

// ServerID is server ID controller.
type ServerID struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerID) ResourceName() string {
	return "server"
}

// Response creates a new response DTO.
func (c *ServerID) Response() base.GetResponseInterface {
	return new(dto.GetServerResponse)
}

// Service returns the service.
func (c *ServerID) Service() base.CRUDServiceInterface {
	return serverService
}
