package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// SwitchSlot saves switch slot info.
type SwitchSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	Index        int                `gorm:"column:Index"`
	Inserted     bool               `gorm:"column:Inserted"`
	SwitchURL    string             `gorm:"column:SwitchURL"`
}

// TableName will set the table name.
func (SwitchSlot) TableName() string {
	return "SwitchSlot"
}

// ToModel converts entity to model.
func (e SwitchSlot) ToModel() *model.SwitchSlot {
	m := model.SwitchSlot{}
	m.Index = e.Index
	m.Inserted = e.Inserted
	return &m
}

// Load loads the model to entity.
func (e SwitchSlot) Load(m *model.SwitchSlot) {
	e.Index = m.Index
	e.Inserted = m.Inserted
}
