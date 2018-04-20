package model

import (
	"promise/base"
)

// ServerGroup is the model of servergroup.
type ServerGroup struct {
	base.Model
	Name        string
	Description string
}

// GetDebugName return the debug name the model.
func (m *ServerGroup) GetDebugName() string {
	return m.Name
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *ServerGroup) GetValueForDuplicationCheck() string {
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
