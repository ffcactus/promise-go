package entity

import (
	"promise/common/app"
	"promise/server/object/model"
	"time"
)

// Group is the entity of server group.
type Group struct {
	Entity
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

// ToModel change the entity to model.
func (e *Group) ToModel() *model.Group {
	ret := new(model.Group)
	ret.ID = e.ID
	ret.URI = app.RootURL + "/servergroup/" + e.ID
	ret.Name = e.Name
	ret.Description = e.Description
	return ret
}

// Load will load the model to entity.
func (e *Group) Load(m *model.Group) {
	e.ID = m.ID
	e.Name = m.Name
	e.Description = m.Description
}
