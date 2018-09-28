package mm920

import (
	"promise/enclosure/object/model"
	log "github.com/sirupsen/logrus"
)

// Client implements EnclosureClient interface.
type Client struct {
	Address string
	Username string
	Password string
}

// String returns the debug info of the client.
func (c Client) String() string {
	return Address
}

// BladeSlot returns the blade slot info.
func (c Client) BladeSlot() ([]model.BladeSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get blade slot.")
	return nil, nil
}

// SwitchSlot returns the switch ade slot info.
func (c Client) SwitchSlot() ([]model.SwitchSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get switch slot.")
	return nil, nil
}

// FanSlot returns the fan slot info.
func (c Client) FanSlot() ([]model.FanSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get fan slot.")
	return nil, nil
}

// PowerSlot returns the power slot info.
func (c Client) PowerSlot() ([]model.PowerSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get power slot.")
	return nil, nil
}

// ManagerSlot returns the manager slot info.
func (c Client) ManagerSlot() ([]model.ManagerSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get manager slot.")
	return nil, nil
}

// ApplianceSlot returns the manager slot info.
func (c Client) ApplianceSlot() ([]model.ApplianceSlot, Error) {
	log.WithFields(log.Fields{"client": c}).Info("Client get appliance slot.")
	return nil, nil
}
