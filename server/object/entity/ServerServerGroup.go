package entity

import (
	"promise/server/object/model"
	"time"
)

// ServerServerGroup is the entity of servergroup.
type ServerServerGroup struct {
	Entity
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ServerID      string
	ServerGroupID string
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
