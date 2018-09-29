package enclosure

import (
	// log "github.com/sirupsen/logrus"
	"fmt"
	"promise/base"
	"promise/enclosure/client/enclosure/mm920"
	"promise/enclosure/client/enclosure/mock"
	"promise/enclosure/object/model"
)

// ClientErrorImpl holds the error info.
// ClientErrorImpl implements Error interface.
type ClientErrorImpl struct {
	status          int
	body            []byte
	connectionError bool
	requestError	error
	timeout         bool
	loginFailure    bool
}

// String returns the debug info for the client error.
func (e ClientErrorImpl) String() string {
	if e.requestError != nil {
		return fmt.Sprintf("%v", e.requestError)
	}
	if timeout {
		return "timeout"
	}
	if connectionError {
		return fmt.Sprintf("(connection error, status = %d)", e.status)
	}
	if loginFailure {
		return fmt.Sprintf("(login failure, status = %d)", e.status)
	}
	return fmt.Sprintf("status = %d", e.status)
}

// Client is the client interface for enclosure device.
type Client interface {
	Ready() bool
	Claim() base.ClientError
	Unclaim() base.ClientError
	DeviceIdentity() (*base.DeviceIdentity, base.ClientError)
	BladeSlot() ([]model.BladeSlot, base.ClientError)
	SwitchSlot() ([]model.SwitchSlot, base.ClientError)
	FanSlot() ([]model.FanSlot, base.ClientError)
	PowerSlot() ([]model.PowerSlot, base.ClientError)
	ManagerSlot() ([]model.ManagerSlot, base.ClientError)
	ApplianceSlot() ([]model.ApplianceSlot, base.ClientError)
}

// NewClient creates a enclosure client by enclosure.
func NewClient(enclosure *model.Enclosure) Client {
	switch enclosure.Type {
	case model.EnclosureTypeMock:
		return mock.NewClient(enclosure)
	case model.EnclosureTypeE9000:
		return mm920.NewClient(enclosure)
	default:
		return nil
	}
}

func getCredential(enclosure model.Enclosure) (string, string, base.ServiceError) {
	return enclosure.Credential.Username, enclosure.Credential.Password, nil
}
