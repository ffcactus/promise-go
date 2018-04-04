package model

import (
	"promise/common/object/model"
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
	model.PromiseModel
	Name        string
	Description *string
	Ranges      []IPv4Range
	SubnetMask  *string
	Gateway     *string
	Domain      *string
	DNSServers  *[]string
}
