package controller

import (
	"promise/base"
	"promise/pool/object/dto"
	"promise/pool/service"
)

var (
	allocate = base.ActionInfo{
		Name: "allocate",
		Request: &dto.AllocateIPv4Request{
			ActionRequest: base.ActionRequest{
				TemplateImpl: new(dto.AllocateIPv4Request),
			},
		},
		Service: &base.ActionService{
			TemplateImpl: new(service.Allocate),
		},
	}
	free = base.ActionInfo{
		Name: "free",
		Request: &dto.FreeIPv4Request{
			ActionRequest: base.ActionRequest{
				TemplateImpl: new(dto.FreeIPv4Request),
			},
		},
		Service: &base.ActionService{
			TemplateImpl: new(service.Free),
		},
	}

	actionInfo = []base.ActionInfo{allocate, free}
)

// IPv4ActionController is implements ActionControllerTemplateInterface.
type IPv4ActionController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *IPv4ActionController) GetResourceName() string {
	return "ipv4"
}

// GetActionInfo returns the name this controller handle of.
func (c *IPv4ActionController) GetActionInfo() []base.ActionInfo {
	return actionInfo
}
