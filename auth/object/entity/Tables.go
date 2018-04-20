package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Account", Info: new(Account)},
		{Name: "Session", Info: new(Session)},
	}
)
