package model

import (
	"promise/common/object/model"
)

// IPv4Range is a IPv4 range.
type IPv4Range struct {
	Start string
	End   string
}

// IPv4Pool is the model.
type IPv4Pool struct {
	model.PromiseModel
	Name        string
	Description *string
	Ranges      []IPv4Range
	SubnetMask  string
	Gateway     string
	Domain      string
	DNSServers  []string
}
