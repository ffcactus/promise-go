package db

import (
	"promise/server/object/model"
)

// ServerGroupDBInterface is the DB interface
type ServerGroupDBInterface interface {
	GetServerGroupByName(name string) *model.ServerGroup
	PostServerGroup(s *model.ServerGroup) (*model.ServerGroup, bool, error)
	GetServerGroup(id string) *model.ServerGroup
	GetServerGroupCollection(start int64, count int64, filter string) (*model.ServerGroupCollection, error)
	DeleteServerGroup(id string) (*model.ServerGroup, error)
	DeleteServerGroupCollection() error
}
