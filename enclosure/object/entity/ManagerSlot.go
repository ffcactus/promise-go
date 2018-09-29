package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// ManagerSlot saves switch slot info.
type ManagerSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	Index        int                `gorm:"column:Index"`
	Inserted     bool               `gorm:"column:Inserted"`
}

// TableName will set the table name.
func (ManagerSlot) TableName() string {
	return "ManagerSlot"
}

// ToModel converts entity to model.
func (e ManagerSlot) ToModel() *model.ManagerSlot {
	m := model.ManagerSlot{}
	m.Index = e.Index
	m.Inserted = e.Inserted
	return &m
}

// Load loads the model to entity.
func (e ManagerSlot) Load(m *model.ManagerSlot) {
	e.Index = m.Index
	e.Inserted = m.Inserted
}
