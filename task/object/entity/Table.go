package entity

import (
	commonDB "promise/common/db"
)

var (
	// Tables The tables used by this project.
	Tables = []commonDB.TableInfo{
		{"Task", new(Task)},
	}
)
