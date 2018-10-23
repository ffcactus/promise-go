package dto

import (
	"promise/enclosure/object/model"
)

// ServerSlot represents the blade slot info.
type ServerSlot struct {
	model.ServerSlotCommon
}

// Load will loads info from model.
func (dto *ServerSlot) Load(m *model.ServerSlot) {
	dto.ServerSlotCommon = m.ServerSlotCommon
}
