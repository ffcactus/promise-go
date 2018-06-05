package model

import (
	"promise/base"
)

// AdapterConfig is the model.
type AdapterConfig struct {
	base.Model
	Name string
}

// DebugInfo return the debug name the model.
func (m *AdapterConfig) DebugInfo() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *AdapterConfig) ValueForDuplicationCheck() string {
	return ""
}

// AdapterConfigCollectionMember is the member in collection.
type AdapterConfigCollectionMember struct {
	base.CollectionMemberModel
	Name string
}

// AdapterConfigCollection is the model of collection.
type AdapterConfigCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *AdapterConfigCollection) NewModelMember() interface{} {
	return new(AdapterConfigCollectionMember)
}
