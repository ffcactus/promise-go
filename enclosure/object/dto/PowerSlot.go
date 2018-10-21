package dto

import (
	"promise/enclosure/object/model"
)

// PowerSlot represents the blade slot info.
type PowerSlot struct {
	model.PowerSlotCommon
}

// Load will loads info from model.
func (dto *PowerSlot) Load(m *model.PowerSlot) {
	dto.PowerSlotCommon = m.PowerSlotCommon
}