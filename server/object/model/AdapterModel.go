package model

import (
	"promise/base"
)

// RAIDCapability describe the capability of an FCoE adapter.
type RAIDCapability struct {
	Version int
}

// EthernetCapability describe the capability of an FCoE adapter.
type EthernetCapability struct {
	Version int
}

// FCoECapability describe the capability of an FCoE adapter.
type FCoECapability struct {
	Version int
}

// AdapterCapability describe the capability of an adapter.
type AdapterCapability struct {
	Version  int
	RAID     *RAIDCapability
	Ethernet *EthernetCapability
	FCoE     *FCoECapability
}

// AdapterModel is the model.
type AdapterModel struct {
	base.Model
	Name       string
	Type       string
	Capability AdapterCapability
}

// String return the debug name the model.
func (m AdapterModel) String() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *AdapterModel) ValueForDuplicationCheck() string {
	return m.Name
}

// AdapterModelCollectionMember is the member in collection.
type AdapterModelCollectionMember struct {
	base.CollectionMemberModel
	Name string
	Type string
}

// AdapterModelCollection is the model of collection.
type AdapterModelCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *AdapterModelCollection) NewModelMember() interface{} {
	return new(AdapterModelCollectionMember)
}
