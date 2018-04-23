package controller

import (
	"promise/base"
	"promise/server/service"
)

var (
	discover = base.ActionInfo{
		Name:    "discover",
		Type:    base.ActionTypeAsych,
		Request: new(dto.DiscoverServerRequest),
		Service: new(service.Discover),
	}

	actionInfo = []base.ActionInfo{discover}
)

// ServerDiscover Server action controller
type ServerDiscover struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerDiscover) ResourceName() string {
	return "server"
}

// ActionInfo returns the name this controller handle of.
func (c *ServerDiscover) ActionInfo() []base.ActionInfo {
	return actionInfo
}
