package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// PowerSlot saves switch slot info.
type PowerSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	Index        int                `gorm:"column:Index"`
	Inserted     bool               `gorm:"column:Inserted"`
}

// TableName will set the table name.
func (PowerSlot) TableName() string {
	return "PowerSlot"
}

// ToModel converts entity to model.
func (e PowerSlot) ToModel() *model.PowerSlot {
	m := model.PowerSlot{}
	m.Index = e.Index
	m.Inserted = e.Inserted
	return &m
}

// Load loads the model to entity.
func (e PowerSlot) Load(m *model.PowerSlot) {
	e.Index = m.Index
	e.Inserted = m.Inserted
}
