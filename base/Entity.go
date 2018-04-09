package base

import (
	"time"
)

// EntityInterface is the interface of a Promise entity should have.
type EntityInterface interface {
	ToModel() ModelInterface
	ToMember() MemberInterface
	Load(ModelInterface)
}

// EntityRefType is the ID type of Entity.
type EntityRefType string

// Entity is the entity used in Promise project.
type Entity struct {
	ID        string    `gorm:"column:ID;primary_key"`
	Category  string    `gorm:"column:Category"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

// ToModel change the entity to model.
func (e *Entity) ToModel() ModelInterface {
	var m Model

	m.ID = e.ID
	m.Category = e.Category
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	return &m
}

// ToMember change the entity to a collection member.
func (e *Entity) ToMember() MemberInterface {
	var m Member
	
	m.ID = e.ID
	m.Category = e.Category
	return &m
}