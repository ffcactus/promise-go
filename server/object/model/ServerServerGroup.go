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

// String return the debug name the model.
func (m ServerServerGroup) String() string {
	return m.ServerID + " " + m.ServerGroupID
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *ServerServerGroup) ValueForDuplicationCheck() string {
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
	return new(ServerServerGroupCollectionMember)
}
