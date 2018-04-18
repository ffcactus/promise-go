package model

import (
	"promise/base"
)

// IPv4Address shows the usage of an IPv4 address.
type IPv4Address struct {
	Key       string
	Address   string
	Allocated bool
}

// IPv4Range is a IPv4 range.
type IPv4Range struct {
	Start       string
	End         string
	Addresses   []IPv4Address
	Total       uint32
	Free        uint32
	Allocatable uint32
}

// IPv4Pool is the model.
type IPv4Pool struct {
	base.Model
	Name        string
	Description *string
	Ranges      []IPv4Range
	SubnetMask  *string
	Gateway     *string
	Domain      *string
	DNSServers  *[]string
}

// GetDebugName return the debug name the model.
func (m *IPv4Pool) GetDebugName() string {
	return m.Name
}

// GetValueForDuplicationCheck return the value for duplication check.
func (m *IPv4Pool) GetValueForDuplicationCheck() string {
	return m.Name
}

// IPv4PoolCollectionMember is the member in collection.
type IPv4PoolCollectionMember struct {
	base.CollectionMemberModel
	Name string
}

// IPv4PoolCollection is the model of collection.
type IPv4PoolCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *IPv4PoolCollection) NewModelMember() interface{} {
	return new(IPv4PoolCollectionMember)
}
