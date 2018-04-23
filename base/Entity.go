package base

import (
	"time"
)

// EntityTemplateInterface is the interface that a concrete entity should have.
type EntityTemplateInterface interface {
	TableName() string
	PropertyNameForDuplicationCheck() string
	DebugInfo() string
	Preload() []string
	Association() []interface{}
	FilterNameList() []string
}

// EntityInterface is the interface of a Promise entity should have.
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

// DebugInfo return the debug name of this entity.
func (e *Entity) DebugInfo() string {
	return e.TemplateImpl.DebugInfo()
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Entity) PropertyNameForDuplicationCheck() string {
	return e.TemplateImpl.PropertyNameForDuplicationCheck()
}

// Preload return the property names that need to be preload.
func (e *Entity) Preload() []string {
	return e.TemplateImpl.Preload()
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *Entity) Association() []interface{} {
	return e.TemplateImpl.Association()
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *Entity) Tables() []interface{} {
	return e.TemplateImpl.Tables()
}

// FilterNameList return all the property name that can be used in filter.
func (e *Entity) FilterNameList() []string {
	return e.TemplateImpl.FilterNameList()
}

// Load will load info from model.
func (e *Entity) Load(i ModelInterface) error {
	return e.TemplateImpl.Load(i)
}

// ToModel convert the entity to model.
func (e *Entity) ToModel() ModelInterface {
	return e.TemplateImpl.ToModel()
}

// ToCollectionMember convert the entity to member.
func (e *Entity) ToCollectionMember() CollectionMemberModelInterface {
	return e.TemplateImpl.ToCollectionMember()
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

// ElementEntityInterface is the interface that ElementEntity should have.
type ElementEntityInterface interface {
	TableName() string
	ToModel() interface{}
	Load(p interface{}) error
}

// ElementEntity represents an element in an array in entity.
type ElementEntity struct {
	TemplateImpl ElementEntityInterface `gorm:"-" json:"-"`
	ID           uint64                 `gorm:"column:ID;primary_key"`
}

// TableName will set the table name.
func (e *ElementEntity) TableName() string {
	return e.TemplateImpl.TableName()
}

// ToModel convert the entity to model.
func (e *ElementEntity) ToModel() interface{} {
	return e.TemplateImpl.ToModel()
}

// Load will load info from model.
func (e *ElementEntity) Load(p interface{}) error {
	return e.TemplateImpl.Load(p)
}
