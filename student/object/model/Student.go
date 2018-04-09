package model

import (
	"promise/base"
)
// Student is the model used in student project.
type Student struct {
	base.Model
	Name string
	Age int
}

