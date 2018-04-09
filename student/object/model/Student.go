package model

import (
	"promise/base"
)

// Student is the model used in student project.
type Student struct {
	base.Model
	Name string
	Age  int
}

// GetDebugName return the debug name the model.
func (m *Student) GetDebugName() string {
	return m.Name
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *Student) GetValueForDuplicationCheck() string {
	return m.Name
}
