package base

import (
	"time"
)

// EntityTemplateInterface is the interface that a concrete entity should have.
type EntityTemplateInterface interface {
	GetPropertyNameForDuplicationCheck() string
	GetDebugName() string
	GetPreload() []string
	GetAssociation() []interface{}
}

// EntityInterface is the interface of a Promise entity should have.
type EntityInterface interface {
	GetPropertyNameForDuplicationCheck() string
	GetDebugName() string
	GetPreload() []string
	GetAssociation() []interface{}
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
	TemplateImpl EntityInterface
	ID        string    `gorm:"column:ID;primary_key"`
	Category  string    `gorm:"column:Category"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
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
