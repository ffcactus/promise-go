package model

import (
	"promise/base"
)

// ServerServerGroup is the model of server-servergroup.
type ServerServerGroup struct {
	base.Model
	ServerID      string
	ServerGroupID string
}

// GetDebugName return the debug name the model.
func (m *ServerServerGroup) GetDebugName() string {
	return m.ServerID + " " + m.ServerGroupID
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *ServerServerGroup) GetValueForDuplicationCheck() string {
	return ""
}

// ServerServerGroupCollectionMember is the member in collection.
type ServerServerGroupCollectionMember struct {
	base.CollectionMemberModel
	ServerID      string
	ServerGroupID string
}

// ServerServerGroupCollection is the model of collection.
type ServerServerGroupCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *ServerServerGroupCollection) NewModelMember() interface{} {
	return new(ServerCollectionMember)
}
