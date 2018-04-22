package controller

import (
	"promise/base"
	"promise/pool/object/dto"
	"promise/pool/service"
)

var (
	allocate = base.ActionInfo{
		Name:    "allocate",
		Type:    base.ActionTypeSych,
		Request: new(dto.AllocateIPv4Request),
		Service: new(service.Allocate),
	}
	free = base.ActionInfo{
		Name:    "free",
		Type:    base.ActionTypeSych,
		Request: new(dto.FreeIPv4Request),
		Service: new(service.Free),
	}

	actionInfo = []base.ActionInfo{allocate, free}
)

// IPv4ActionController is implements ActionControllerTemplateInterface.
type IPv4ActionController struct {
}

// ResourceName returns the name this controller handle of.
func (c *IPv4ActionController) ResourceName() string {
	return "ipv4"
}

// ActionInfo returns the name this controller handle of.
func (c *IPv4ActionController) ActionInfo() []base.ActionInfo {
	return actionInfo
}
