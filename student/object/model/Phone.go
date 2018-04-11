package model

import (
	"promise/base"
)

// Phone is the phone model.
type Phone struct {
	base.ModelMember
	Number string
}