package mm920

import (
	"promise/enclosure/object/model"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"net/http"
	"crypto/tls"
)

// Client implements EnclosureClient interface.
type Client struct {
	_client *http.Client
	PreferredAddress string
	Addresses []string
	Username string
	Password string
}

// NewClient creates a client for enclosure.
func NewClient(enclosure *model.Enclosure) *Client {
	client := Client{
		_client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
	}
	// TODO or get username and password from keystore.
	if enclosure.Credential.URL == "" {
		client.Username = enclosure.Credential.Username
		client.Password = enclosure.Credential.Password
	}
	if len(enclosure.Addresses) > 0 {
		client.PreferredAddress = enclosure.Addresses[0]
		client.Addresses = enclosure.Addresses
	}
	return &client
}

// String returns the debug info of the client.
func (c Client) String() string {
	return c.PreferredAddress
}

// DeviceIdentity returns the device identity.
func (c Client) DeviceIdentity() (*base.DeviceIdentity, base.ClientError) {

	return nil, nil
}

// BladeSlot returns the blade slot info.
func (c Client) BladeSlot() ([]model.BladeSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get blade slot.")
	return nil, nil
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
