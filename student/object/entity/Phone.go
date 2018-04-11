package entity

import (
	"promise/base"
	"promise/student/object/model"
)

// Phone is the phone entity.
type Phone struct {
	base.ElementEntity
	StudentRef base.EntityRefType `gorm:"column:StudentRef"`
	Number     string             `gorm:"column:Number"`
}

// TableName will set the table name.
func (e *Phone) TableName() string {
	return "Phone"
}

// Load will load data from model.
// Note that the ID is not load.
func (e *Phone) Load(m *model.Phone) {
	e.Number = m.Number
}

// ToModel will convert entity to model.
func (e *Phone) ToModel() *model.Phone {
	m := model.Phone{}
	m.ID = e.ID
	m.Number = e.Number
	return &m
}
