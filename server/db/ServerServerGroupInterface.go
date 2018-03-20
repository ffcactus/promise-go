package db

import (
	"promise/server/object/model"
)

// ServerServerGroupInterface is the server-servergroup DB interface.
type ServerServerGroupInterface interface {
	PostServerServerGroup(sg *model.ServerServerGroup) (*model.ServerServerGroup, bool, error)
	GetServerServerGroup(id string) *model.ServerServerGroup
	GetServerServerGroupCollection(start int, count int, filter string) (*model.ServerServerGroupCollection, error)
	DeleteServerServerGroupCollection() error
}
