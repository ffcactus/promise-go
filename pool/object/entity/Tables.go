package entity

import (
	commonDB "promise/common/db"
)

var (
	// Tables The tables used by this project.
	Tables = []commonDB.TableInfo{
		{"IPv4Address", new(IPv4Address)},
		{"IPv4Range", new(IPv4Range)},
		{"IPv4Pool", new(IPv4Pool)},
	}
)
