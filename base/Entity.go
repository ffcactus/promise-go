package base

import (
	"time"
)

// EntityTemplateInterface is the interface that a concrete entity should have.
type EntityTemplateInterface interface {
	TableName() string
	GetPropertyNameForDuplicationCheck() string
	GetDebugName() string
	GetPreload() []string
	GetAssociation() []interface{}
}

// EntityInterface is the interface of a Promise entity should have.
type EntityInterface interface {
	TableName() string
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
	TemplateImpl EntityInterface `gorm:"-" json:"-"`
	ID           string          `gorm:"column:ID;primary_key"`
	Category     string          `gorm:"column:Category"`
	CreatedAt    time.Time       `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time       `gorm:"column:UpdatedAt"`
}

// TableName will set the table name.
func (e *Entity) TableName() string {
	return e.TemplateImpl.TableName()
}

// GetID return the ID.
func (e *Entity) GetID() string {
	return e.ID
}

// SetID set the ID.
func (e *Entity) SetID(id string) {
	e.ID = id
}

// GetDebugName return the debug name of this entity.
func (e *Entity) GetDebugName() string {
	return e.TemplateImpl.GetDebugName()
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Entity) GetPropertyNameForDuplicationCheck() string {
	return e.TemplateImpl.GetPropertyNameForDuplicationCheck()
}

// GetPreload return the property names that need to be preload.
func (e *Entity) GetPreload() []string {
	return e.TemplateImpl.GetPreload()
}

// GetAssociation return all the association address that need to delete.
func (e *Entity) GetAssociation() []interface{} {
	return e.TemplateImpl.GetAssociation()
}

// Load will load info from model.
func (e *Entity) Load(i ModelInterface) error {
	return e.TemplateImpl.Load(i)
}

// ToModel convert the entity to model.
func (e *Entity) ToModel() ModelInterface {
	return e.TemplateImpl.ToModel()
}

// ToMember convert the entity to member.
func (e *Entity) ToMember() MemberInterface {
	return e.TemplateImpl.ToMember()
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
	m.Category = e.Category
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}
