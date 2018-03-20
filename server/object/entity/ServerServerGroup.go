package entity

import (
	"promise/common/object/entity"
	"promise/server/object/model"
)

// ServerServerGroup is the entity of servergroup.
type ServerServerGroup struct {
	entity.PromiseEntity
	ServerID      string `gorm:"column:ServerID"`
	ServerGroupID string `gorm:"column:ServerGroupID"`
}

// TableName will set the table name.
func (ServerServerGroup) TableName() string {
	return "ServerServerGroup"
}

// ToModel change the entity to model.
func (e *ServerServerGroup) ToModel() *model.ServerServerGroup {
	ret := new(model.ServerServerGroup)
	ret.ID = e.ID
	ret.ServerID = e.ServerID
	ret.ServerGroupID = e.ServerGroupID
	return ret
}

// Load will load the model to entity.
func (e *ServerServerGroup) Load(m *model.ServerServerGroup) {
	e.ID = m.ID
	e.ServerID = m.ServerID
	e.ServerGroupID = m.ServerGroupID
}
