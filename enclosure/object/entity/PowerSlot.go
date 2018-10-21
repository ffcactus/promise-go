package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// PowerSlot saves switch slot info.
type PowerSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	model.PowerSlotCommon
}

// TableName will set the table name.
func (PowerSlot) TableName() string {
	return "PowerSlot"
}

// ToModel converts entity to model.
func (e PowerSlot) ToModel() *model.PowerSlot {
	m := model.PowerSlot{}
	m.PowerSlotCommon = e.PowerSlotCommon
	return &m
}

// Load loads the model to entity.
func (e *PowerSlot) Load(m *model.PowerSlot) {
	e.PowerSlotCommon = m.PowerSlotCommon
}
