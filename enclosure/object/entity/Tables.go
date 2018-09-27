package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Enclosure", Info: new(Enclosure)},
	}
)