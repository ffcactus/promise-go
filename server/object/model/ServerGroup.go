package model

import (
	"promise/base"
)

// ServerGroup is the model.
type ServerGroup struct {
	base.Model
	Name        string
	Description string
}

// DebugInfo return the debug name the model.
func (m *ServerGroup) DebugInfo() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *ServerGroup) ValueForDuplicationCheck() string {
	return m.Name
}

// ServerGroupCollectionMember is the member in collection.
type ServerGroupCollectionMember struct {
	base.CollectionMemberModel
	Name string
}

// ServerGroupCollection is the model of collection.
type ServerGroupCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *ServerGroupCollection) NewModelMember() interface{} {
	return new(ServerGroupCollectionMember)
}
