package db

import (
	"promise/server/object/model"
)

// GroupDBInterface is the DB interface
type GroupDBInterface interface {
	GetGroupByName(name string) *model.Group
	PostGroup(s *model.Group) (*model.Group, bool, error)
	GetGroup(id string) *model.Group
	GetGroupCollection(start int, count int) (*model.GroupCollection, error)
	DeleteGroup(id string) (bool, error)
	DeleteGroupCollection() error
}
