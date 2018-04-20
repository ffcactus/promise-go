package model

import (
	"promise/common/object/model"
)

// ServerGroup is the model of servergroup.
type ServerGroup struct {
	model.PromiseModel
	Name        string
	Description string
}
