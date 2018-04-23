package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

var (
	discover = base.ActionInfo{
		Name:    "discover",
		Type:    base.ActionTypeAsych,
		Request: new(dto.DiscoverServerRequest),
		Service: new(service.Discover),
	}

	serverRootActionInfo = []base.ActionInfo{discover}
)

// ServerRootAction is the service for actions on server root.
type ServerRootAction struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerRootAction) ResourceName() string {
	return "server"
}

// ActionInfo returns the name this controller handle of.
func (c *ServerRootAction) ActionInfo() []base.ActionInfo {
	return serverRootActionInfo
}
