package db

import (
	"promise/server/object/model"
)

// ServerServerGroupInterface is the server-servergroup DB interface.
type ServerServerGroupInterface interface {
	PostServerServerGroup(sg *model.ServerServerGroup) (*model.ServerServerGroup, bool, error)
	GetServerServerGroup(id string) *model.ServerServerGroup
	GetServerServerGroupCollection(start int64, count int64, filter string) (*model.ServerServerGroupCollection, error)
	DeleteServerServerGroup(id string) (*model.ServerServerGroup, error)
	DeleteServerServerGroupCollection() error
}
