package model

import (
	"promise/common/object/model"
)

// ServerServerGroup is the model of server-servergroup.
type ServerServerGroup struct {
	model.PromiseModel
	ServerID      string
	ServerGroupID string
}
