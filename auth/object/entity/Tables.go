package entity

import (
	commonDB "promise/common/db"
)

var (
	// Tables The tables used by this project.
	Tables = []commonDB.TableInfo{
		{Name: "Account", Info: new(Account)},
		{Name: "Session", Info: new(Session)},
	}
)
