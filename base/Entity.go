package base

import (
	"time"
)

// EntityInterface is the interface that an Entity should have.
type EntityInterface interface {
	TableName() string
	PropertyNameForDuplicationCheck() string
	DebugInfo() string
	Preload() []string
	Association() []interface{}
	Tables() []interface{}
	FilterNameList() []string
	ToModel() ModelInterface
	ToCollectionMember() CollectionMemberModelInterface
	Load(ModelInterface) error
	GetID() string
	SetID(string)
}

// EntityRefType is the ID type of Entity.
type EntityRefType string

// Entity is the entity used in Promise project.
type Entity struct {
	ID           string          `gorm:"column:ID;primary_key"`
	Category     string          `gorm:"column:Category"`
	CreatedAt    time.Time       `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time       `gorm:"column:UpdatedAt"`
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
func EntityToMember(e *Entity, m *CollectionMemberModel) {
	m.ID = e.ID
	m.Category = e.Category
}

// EntityToModel convert entity to model.
func EntityToModel(e *Entity, m *Model) {
	m.ID = e.ID
	m.Category = e.Category
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}

// ElementEntityRefType is the type to define a ref to ArrayElement.
type ElementEntityRefType uint64

// ElementEntity represents an element in an array in entity.
type ElementEntity struct {
	ID           uint64                 `gorm:"column:ID;primary_key"`
}
