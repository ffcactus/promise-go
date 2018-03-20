package entity

import (
	"promise/common/object/entity"
	"promise/server/object/model"
)

// ServerGroup is the entity of servergroup.
type ServerGroup struct {
	entity.PromiseEntity
	Name        string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
}

// TableName will set the table name.
func (ServerGroup) TableName() string {
	return "ServerGroup"
}

// ToModel change the entity to model.
func (e *ServerGroup) ToModel() *model.ServerGroup {
	ret := new(model.ServerGroup)
	ret.PromiseModel = e.PromiseEntity.ToModel()
	ret.Name = e.Name
	ret.Description = e.Description
	return ret
}

// Load will load the model to entity.
func (e *ServerGroup) Load(m *model.ServerGroup) {
	e.ID = m.ID
	e.Name = m.Name
	e.Description = m.Description
}
