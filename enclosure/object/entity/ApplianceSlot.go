package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// ApplianceSlot saves switch slot info.
type ApplianceSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	model.ApplianceSlotCommon
}

// TableName will set the table name.
func (ApplianceSlot) TableName() string {
	return "ApplianceSlot"
}

// ToModel converts entity to model.
func (e ApplianceSlot) ToModel() *model.ApplianceSlot {
	m := model.ApplianceSlot{}
	m.ApplianceSlotCommon = e.ApplianceSlotCommon
	return &m
}

// Load loads the model to entity.
func (e *ApplianceSlot) Load(m *model.ApplianceSlot) {
	e.ApplianceSlotCommon = m.ApplianceSlotCommon
}
