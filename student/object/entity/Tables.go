package entity

import (
	"promise/apps"
)

var (
	// Tables The tables used by this project.
	Tables = []apps.TableInfo{
		{Name: "Student", Info: new(Student)},
	}
)
