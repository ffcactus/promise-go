package mm920

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// Client is MM920's client.
type Client struct {
	base.Client
}

// NewClient creates a client for enclosure.
func NewClient(enclosure *model.Enclosure) *Client {
	client := Client{}
	client.Protocol = "https"
	// TODO or get username and password from keystore.
	if enclosure.Credential.URL == "" {
		client.Username = enclosure.Credential.Username
		client.Password = enclosure.Credential.Password
	}
	if len(enclosure.Addresses) > 0 {
		client.CurrentAddress = enclosure.Addresses[0]
		client.Addresses = enclosure.Addresses
	}
	return &client
}

// Ready returns if the enclosure is ready.
func (c Client) Ready() bool {
	return false
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
	var (
		identity         base.DeviceIdentity
		redfishV1        GetRedfishV1Response
		chassisEnclosure GetChassisEnclosureResponse
	)
	// Get UUID.
	if err := c.Get("/redfish/v1", &redfishV1); err != nil {
		return nil, err
	}
	identity.UUID = redfishV1.UUID
	// Get SerialNumber and PartNumber
	if err := c.Get("/redfish/v1/Chassis/Enclosure", &chassisEnclosure); err != nil {
		return nil, err
	}
	identity.SerialNumber = chassisEnclosure.SerialNumber
	identity.PartNumber = chassisEnclosure.PartNumber

	return &identity, nil
}

// ServerSlot returns the blade slot info.
func (c Client) ServerSlot() ([]model.ServerSlot, base.ClientError) {
	var (
		slots    []model.ServerSlot
		inserted int
	)
	for i := 1; i < 16; i++ {
		chassis := GetBladeChassisResponse{}
		slot := model.ServerSlot{}
		if err := c.Get(fmt.Sprintf("/redfish/v1/Chassis/Blade%d", i), &chassis); err != nil {
			return nil, err
		}
		slot.Index = i
		slot.Inserted = (chassis.Status.State == "Enabled")
		if slot.Inserted {
			inserted++
		}
		slot.ProductName = chassis.Model
		slot.SerialNumber = chassis.SerialNumber
		slot.Height = chassis.Oem.Huawei.Height / 2
		slot.Width = chassis.Oem.Huawei.Width
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"inserted": inserted, "client": c}).Info("Client get blade slot.")
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
