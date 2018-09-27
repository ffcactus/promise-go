package mm920

import (
	"promise/enclosure/object/model"
)

// Client implements EnclosureClient interface.
type Client struct {
}

// BladeSlot returns the blade slot info.
func (c Client) BladeSlot() ([]model.BladeSlot, error) {
	return nil, nil
}

// SwitchSlot returns the switch ade slot info.
func (c Client) SwitchSlot() ([]model.SwitchSlot, error) {
	return nil, nil
}

// FanSlot returns the fan slot info.
func (c Client) FanSlot() ([]model.FanSlot, error) {
	return nil, nil
}

// PowerSlot returns the power slot info.
func (c Client) PowerSlot() ([]model.PowerSlot, error) {
	return nil, nil
}

// ManagerSlot returns the manager slot info.
func (c Client) ManagerSlot() ([]model.ManagerSlot, error) {
	return nil, nil
}

// ApplianceSlot returns the manager slot info.
func (c Client) ApplianceSlot() ([]model.ApplianceSlot, error) {
	return nil, nil
}
