package dto

import (
	"promise/enclosure/object/model"
)

// FanSlot represents the blade slot info.
type FanSlot struct {
	model.FanSlotCommon
}

// Load will loads info from model.
func (dto *FanSlot) Load(m *model.FanSlot) {
	dto.FanSlotCommon = m.FanSlotCommon
}
