package model

import (
	"promise/base"
)

const (
	// EnclosureTypeMock is the enum represents mock enclosure.
	EnclosureTypeMock = "Mock"
	// EnclosureTypeE9000 is the enum represents E9000 enclosure.
	EnclosureTypeE9000 = "E9000"
)

const (
	// StateAdded means enclosure is just added.
	StateAdded = "Added"
	// StateReady means enclosure is ready to take an action.
	StateReady = "Ready"
	// StateLocked means enclosure is under an action.
	StateLocked = "Locked"
	// StateUnmanaged means enclosure is unmanaged.
	StateUnmanaged = "Unmanaged"

	// StateReasonAuto means enclosure change to the state automatically.
	StateReasonAuto = "Auto"

	// HealthOK means enclosure has no alarm.
	HealthOK = "OK"
	// HealthWarning means enclosure has warning alarm.
	HealthWarning = "Warning"
	// HealthCritical means enclosure has critical alarm.
	HealthCritical = "Critical"
)

// EnclosureLockable check if the enclosure's state can turn to be Locked.
func EnclosureLockable(state string) bool {
	var ret = false
	switch state {
	case StateAdded, StateReady, StateUnmanaged:
		ret = true
	default:
		ret = false
	}
	return ret
}

// Enclosure is the model of enclosure.
type Enclosure struct {
	base.Model
	base.DeviceIdentity
	Name           string
	Description    string
	Type           string
	State          string
	Health         string
	Addresses      []string // The addresses that can be used to connect to enclosure.
	Credential     Credential
	ServerSlots    []ServerSlot
	SwitchSlots    []SwitchSlot
	ManagerSlots   []ManagerSlot
	ApplianceSlots []ApplianceSlot
	FanSlots       []FanSlot
	PowerSlots     []PowerSlot
}

// Credential should contains URL that can retrieve the cridentail or username and password
type Credential struct {
	URL      string
	Username string
	Password string
}

// String return the debug name the model.
func (m Enclosure) String() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *Enclosure) ValueForDuplicationCheck() string {
	return m.Name
}
