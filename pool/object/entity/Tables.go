package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "IPv4Address", Info: new(IPv4Address)},
		{Name: "IPv4Range", Info: new(IPv4Range)},
		{Name: "IPv4Pool", Info: new(IPv4Pool)},
	}
)
