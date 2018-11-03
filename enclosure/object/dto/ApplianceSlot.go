package dto

import (
	"promise/enclosure/object/model"
)

// ApplianceSlot represents the blade slot info.
type ApplianceSlot struct {
	model.ApplianceSlotCommon
}

// Load will loads info from model.
func (dto *ApplianceSlot) Load(m *model.ApplianceSlot) {
	dto.ApplianceSlotCommon = m.ApplianceSlotCommon
}
