package controller

import (
	"promise/base"
	"promise/server/service"
)

var (
	refresh = base.ActionInfo{
		Name:    "refresh",
		Type:    base.ActionTypeAsych,
		Request: nil,
		Service: new(service.Refresh),
	}

	serverAction = []base.ActionInfo{refresh}
)

// IPv4ActionController is implements ActionControllerTemplateInterface.
type ServerAction struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerAction) ResourceName() string {
	return "server"
}

// ActionInfo returns the name this controller handle of.
func (c *ServerAction) ActionInfo() []base.ActionInfo {
	return serverAction
}
