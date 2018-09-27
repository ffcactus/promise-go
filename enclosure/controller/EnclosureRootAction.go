package controller

import (
	"promise/base"
	"promise/enclosure/object/dto"
	"promise/enclosure/service"
)

var (
	discover = base.ActionInfo{
		Name:    "discover",
		Type:    base.ActionTypeSych,
		Request: new(dto.DiscoverEnclosureRequest),
		Service: new(service.Discover),
	}

	enclosureRootActionInfo = []base.ActionInfo{discover}
)

// EnclosureRootAction is the service for actions on enclosure root.
type EnclosureRootAction struct {
}

// ResourceName returns the name this controller handle of.
func (c *EnclosureRootAction) ResourceName() string {
	return "enclosure"
}

// ActionInfo returns the actions this controller handle of.
func (c *EnclosureRootAction) ActionInfo() []base.ActionInfo {
	return enclosureRootActionInfo
}
