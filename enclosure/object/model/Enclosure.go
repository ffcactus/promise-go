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

// Enclosure is the model of enclosure.
type Enclosure struct {
	base.Model
	Name          string
	Description   string
	Type          string
	State         string
	Health        string
	Addresses     []string
	BladeSlots    []BladeSlot
	SwitchSlots   []SwitchSlot
	FanSlots      []FanSlot
	PowerSlot     []PowerSlot
	ApplianceSlot []ApplianceSlot
}

// String return the debug name the model.
func (m Enclosure) String() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *Enclosure) ValueForDuplicationCheck() string {
	return m.Name
}
