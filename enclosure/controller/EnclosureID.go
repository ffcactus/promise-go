package controller

import (
	"promise/base"
	"promise/enclosure/object/dto"
)

// EnclosureID is enclosure ID controller.
type EnclosureID struct {
}

// ResourceName returns the name this controller handle of.
func (c *EnclosureID) ResourceName() string {
	return "enclosure"
}

// Response creates a new response DTO.
func (c *EnclosureID) Response() base.GetResponseInterface {
	return new(dto.GetEnclosureResponse)
}

// Service returns the service.
func (c *EnclosureID) Service() base.CRUDServiceInterface {
	return enclosureService
}
