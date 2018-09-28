package client

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/client/mm920"
	"promise/enclosure/client/mock"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
	"fmt"
)

// Error represents the error in client.
type Error Interface {
	Status() int
	Body() []byte
	ConnectionError() bool
	Timeout() bool
	LoginFailure() bool
	String() string
}

// ErrorImpl holds the error info.
// ErrorImpl implements Error interface.
type ErrorImpl {
	status int
	body []byte
	connectionError bool
	timeout bool
	loginFailure bool
}

// String returns the debug info for the client error.
func (e ErrorImpl) String() string {
	fmt.Sprintf("status = %d, timeout = %v, loginFailure = %v", e.status, e.timeout, e.loginFailure)
}

// Client is the client interface for enclosure device.
type Client interface {
	BladeSlot() ([]model.BladeSlot, Error)
	SwitchSlot() ([]model.SwitchSlot, Error)
	FanSlot() ([]model.FanSlot, Error)
	PowerSlot() ([]model.PowerSlot, Error)
	ManagerSlot() ([]model.ManagerSlot, Error)
	ApplianceSlot() ([]model.ApplianceSlot, Error)
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

func getCredential(enclosure model.Enclosure) (username, password string, base.ServiceError) {
	
}
