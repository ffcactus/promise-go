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

// ToModel will create a new model from entity.
func (e *ServerServerGroup) ToModel() *model.ServerServerGroup {
	ret := new(model.ServerServerGroup)
	ret.PromiseModel = e.PromiseEntity.ToModel()
	ret.ServerID = e.ServerID
	ret.ServerGroupID = e.ServerGroupID
	return ret
}

// Load will load data from model.
func (e *ServerServerGroup) Load(m *model.ServerServerGroup) {
	e.PromiseEntity.Load(m.PromiseModel)
	e.ServerID = m.ServerID
	e.ServerGroupID = m.ServerGroupID
}
