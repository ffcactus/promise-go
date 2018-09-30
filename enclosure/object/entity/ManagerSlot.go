package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// ManagerSlot saves switch slot info.
type ManagerSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	model.ManagerSlotCommon
}

// TableName will set the table name.
func (ManagerSlot) TableName() string {
	return "ManagerSlot"
}

// ToModel converts entity to model.
func (e ManagerSlot) ToModel() *model.ManagerSlot {
	m := model.ManagerSlot{}
	m.ManagerSlotCommon = e.ManagerSlotCommon
	return &m
}

// Load loads the model to entity.
func (e *ManagerSlot) Load(m *model.ManagerSlot) {
	e.ManagerSlotCommon = m.ManagerSlotCommon
}
