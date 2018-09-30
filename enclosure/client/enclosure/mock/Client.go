package mock

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// Client implements EnclosureClient interface.
type Client struct {
	Address  string
	Username string
	Password string
}

// NewClient creates a client for enclosure.
func NewClient(enclosure *model.Enclosure) *Client {
	client := Client{}
	if len(enclosure.Addresses) > 0 {
		client.Address = enclosure.Addresses[0]
	}
	return &client
}

// String returns the debug info of the client.
func (c Client) String() string {
	return c.Address
}

// Ready returns if the enclosure is ready.
func (c Client) Ready() bool {
	return true
}

// Claim should make make a flag on the enclosure that indicate it is exclusively managed by this system.
func (c Client) Claim() base.ClientError {
	return nil
}

// Unclaim should remove the flag that indicate the enclosure is managed by this system.
func (c Client) Unclaim() base.ClientError {
	return nil
}

// DeviceIdentity returns the device identity.
func (c Client) DeviceIdentity() (*base.DeviceIdentity, base.ClientError) {
	identity := base.DeviceIdentity{}
	identity.UUID = base.RandUUID()
	identity.SerialNumber = base.RandString(12)
	identity.PartNumber = base.RandString(10)
	log.WithFields(log.Fields{"client": c, "identity": identity}).Info("Client get device identity.")
	return &identity, nil
}

// ServerSlot returns the blade slot info.
func (c Client) ServerSlot() ([]model.ServerSlot, base.ClientError) {
	slots := make([]model.ServerSlot, 0)
	for i:=1; i <= 8; i++ {
		slot := model.ServerSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.ProductName = "CH121 V5"
		slot.SerialNumber = base.RandString(12)
		slot.Height = 1
		slot.Width = 1
		slots = append(slots, slot) 
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get blade slot.")
	return slots, nil
}

// SwitchSlot returns the switch ade slot info.
func (c Client) SwitchSlot() ([]model.SwitchSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get switch slot.")
	return nil, nil
}

// FanSlot returns the fan slot info.
func (c Client) FanSlot() ([]model.FanSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get fan slot.")
	return nil, nil
}

// PowerSlot returns the power slot info.
func (c Client) PowerSlot() ([]model.PowerSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get power slot.")
	return nil, nil
}

// ManagerSlot returns the manager slot info.
func (c Client) ManagerSlot() ([]model.ManagerSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get manager slot.")
	return nil, nil
}

// ApplianceSlot returns the manager slot info.
func (c Client) ApplianceSlot() ([]model.ApplianceSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get appliance slot.")
	return nil, nil
}
