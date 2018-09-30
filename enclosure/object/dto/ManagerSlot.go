package dto

import (
	"promise/enclosure/object/model"
)

// ManagerSlot represents the blade slot info.
type ManagerSlot struct {
	model.ManagerSlotCommon
}

// Load will loads info from model.
func (dto *ManagerSlot) Load(m *model.ManagerSlot) {
	dto.ManagerSlotCommon = m.ManagerSlotCommon
}