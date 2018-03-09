package entity

import (
	"promise/common/app"
	"promise/server/object/model"
	"time"
)

// ServerGroup is the entity of server group.
type ServerGroup struct {
	Entity
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

// ToModel change the entity to model.
func (e *ServerGroup) ToModel() *model.ServerGroup {
	ret := new(model.ServerGroup)
	ret.ID = e.ID
	ret.URI = app.RootURL + "/servergroup/" + e.ID
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
