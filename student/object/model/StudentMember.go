package model

import (
	"promise/base"
)

// StudentMember is member in student collection.
type StudentMember struct {
	base.Member
	Name string
}
