package entity

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// ServerSlot saves switch slot info.
type ServerSlot struct {
	base.ElementEntity
	EnclosureRef base.EntityRefType `gorm:"column:EnclosureRef"`
	model.ServerSlotCommon
}

// TableName will set the table name.
func (ServerSlot) TableName() string {
	return "ServerSlot"
}

// ToModel converts entity to model.
func (e ServerSlot) ToModel() *model.ServerSlot {
	m := model.ServerSlot{}
	m.ServerSlotCommon = e.ServerSlotCommon
	return &m
}

// Load loads the model to entity.
func (e *ServerSlot) Load(m *model.ServerSlot) {
	e.ServerSlotCommon = m.ServerSlotCommon
}
