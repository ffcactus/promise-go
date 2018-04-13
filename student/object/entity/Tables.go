package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Student", Info: new(Student)},
		{Name: "Phone", Info: new(Phone)},
	}
)
