package client

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/client/mm920"
	"promise/enclosure/client/mock"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// Error represents the error in client.
type Error struct {
	error
}

// Client is the client interface for enclosure device.
type Client interface {
	BladeSlot() ([]model.BladeSlot, error)
	SwitchSlot() ([]model.SwitchSlot, error)
	FanSlot() ([]model.FanSlot, error)
	PowerSlot() ([]model.PowerSlot, error)
	ManagerSlot() ([]model.ManagerSlot, error)
	ApplianceSlot() ([]model.ApplianceSlot, error)
}

// NewClient creates a enclosure client by enclosure.
func NewClient(enclosure model.Enclosure) (Client, error) {
	switch enclosure.Type {
	case model.EnclosureTypeMock:
		return nil, nil
	default:
		return nil, nil
	}
}
