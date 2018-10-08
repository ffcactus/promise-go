package controller

import (
	"promise/base"
	"promise/enclosure/object/dto"
	"promise/enclosure/service"
)

var (
	refresh = base.ActionInfo{
		Name:    "refresh",
		Type:    base.ActionTypeAsych,
		Request: new(dto.RefreshEnclosureRequest),
		Service: new(service.Refresh),
	}

	serverAction = []base.ActionInfo{refresh}
)

// EnclosureAction is implements ActionControllerTemplateInterface.
type EnclosureAction struct {
}

// ResourceName returns the name this controller handle of.
func (c *EnclosureAction) ResourceName() string {
	return "enclosure"
}

// ActionInfo returns the name this controller handle of.
func (c *EnclosureAction) ActionInfo() []base.ActionInfo {
	return serverAction
}
