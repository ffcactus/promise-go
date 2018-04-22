package controller

import (
	log "github.com/sirupsen/logrus"
	"promise/server/service"
	"strings"
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

// ServerActionController Server action controller
type ServerActionController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerActionController) ResourceName() string {
	return "server"
}

// ActionInfo returns the name this controller handle of.
func (c *ServerActionController) ActionInfo() []base.ActionInfo {
	return actionInfo
}
