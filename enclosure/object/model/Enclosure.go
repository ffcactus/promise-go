package model

import (
	"promise/base"
)

// Enclosure is the model of enclosure.
type Enclosure struct {
	base.Model
	Name           string
	Description    string
	State          string
	Health         string
	Addresses      []string
}

// String return the debug name the model.
func (m Enclosure) String() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *Enclosure) ValueForDuplicationCheck() string {
	return m.Name
}