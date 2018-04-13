package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Task", Info: new(Task)},
	}
)
