package base

import (
	"time"
)

// EntityInterface is the interface of a Promise entity should have.
type EntityInterface interface {
	GetPropertyNameForDuplicationCheck() string
	GetDebugName() string
	ToModel() ModelInterface
	ToMember() MemberInterface
	Load(ModelInterface) error
	GetID() string
	SetID(string)
}

// EntityRefType is the ID type of Entity.
type EntityRefType string

// Entity is the entity used in Promise project.
type Entity struct {
	// Interface EntityInterface
	ID        string    `gorm:"column:ID;primary_key"`
	Category  string    `gorm:"column:Category"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

// GetID return the ID.
func (e *Entity) GetID() string {
	return e.ID
}

// SetID set the ID.
func (e *Entity) SetID(id string) {
	e.ID = id
}

// EntityLoad load model to entity.
func EntityLoad(e *Entity, m *Model) {
	e.ID = m.ID
	e.Category = m.Category
}

// EntityToMember convert entity to member.
func EntityToMember(e *Entity, m *Member) {
	m.ID = e.ID
	m.Category = e.Category
}

// EntityToModel convert entity to model.
func EntityToModel(e *Entity, m *Model) {
	m.ID = e.ID
	m.Category = e.ID
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}

// // ToModel change the entity to model.
// func (e *Entity) ToModel() ModelInterface {
// 	var m = e.Interface.NewModel()

// 	m.SetID(e.ID)
// 	m.Category = e.Category
// 	m.CreatedAt = e.CreatedAt
// 	m.UpdatedAt = e.UpdatedAt
// 	return m
// }

// // ToMember change the entity to a collection member.
// func (e *Entity) ToMember() MemberInterface {
// 	var m Member

// 	m.ID = e.ID
// 	m.Category = e.Category
// 	return &m
// }
