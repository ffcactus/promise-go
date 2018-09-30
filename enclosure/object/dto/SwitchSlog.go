package dto

import (
	"promise/enclosure/object/model"
)

// SwitchSlot represents the blade slot info.
type SwitchSlot struct {
	model.SwitchSlotCommon
}

// Load will loads info from model.
func (dto *SwitchSlot) Load(m *model.SwitchSlot) {
	dto.SwitchSlotCommon = m.SwitchSlotCommon
}