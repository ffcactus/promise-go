package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// BladeSlot saves switch slot info.
type BladeSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	Index        int                `gorm:"column:Index"`
	Inserted     bool               `gorm:"column:Inserted"`
	ServerURL    string             `gorm:"column:ServerURL"`
}

// TableName will set the table name.
func (BladeSlot) TableName() string {
	return "BladeSlot"
}

// ToModel converts entity to model.
func (e BladeSlot) ToModel() *model.BladeSlot {
	m := model.BladeSlot{}
	m.Index = e.Index
	m.Inserted = e.Inserted
	m.ServerURL = e.ServerURL
	return &m
}

// Load loads the model to entity.
func (e BladeSlot) Load(m *model.BladeSlot) {
	e.Index = m.Index
	e.Inserted = m.Inserted
	e.ServerURL = m.ServerURL
}
