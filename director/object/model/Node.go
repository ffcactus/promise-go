package model

import (
	"promise/base"
)

// Node is the model.
type Node struct {
	base.Model
	Hostname      string
	Status        string
	Availibility  string
	ManagerStatus string
}

// DebugInfo return the debug name the model.
func (m *Node) DebugInfo() string {
	return m.Hostname
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *Node) ValueForDuplicationCheck() string {
	return ""
}

// NodeCollectionMember is the member in collection.
type NodeCollectionMember struct {
	base.CollectionMemberModel
	Hostname      string
	Status        string
	Availibility  string
	ManagerStatus string
}

// NodeCollection is the model of collection.
type NodeCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *NodeCollection) NewModelMember() interface{} {
	return new(NodeCollectionMember)
}
