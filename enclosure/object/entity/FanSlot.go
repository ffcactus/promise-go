package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// FanSlot saves switch slot info.
type FanSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	Index        int                `gorm:"column:Index"`
	Inserted     bool               `gorm:"column:Inserted"`
}

// TableName will set the table name.
func (FanSlot) TableName() string {
	return "FanSlot"
}

// ToModel converts entity to model.
func (e FanSlot) ToModel() *model.FanSlot {
	m := model.FanSlot{}
	m.Index = e.Index
	m.Inserted = e.Inserted
	return &m
}

// Load loads the model to entity.
func (e FanSlot) Load(m *model.FanSlot) {
	e.Index = m.Index
	e.Inserted = m.Inserted
}
