package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// GetEnclosureResponse is the DTO.
type GetEnclosureResponse struct {
	base.GetResponse
	Name        string
	Description string
	State       string
	Health      string
	base.DeviceIdentity
	Addresses      []string
	ServerSlots    []ServerSlot
	SwitchSlots    []SwitchSlot
	ManagerSlots   []ManagerSlot
	ApplianceSlots []ApplianceSlot
	FanSlots       []FanSlot
	PowerSlots     []PowerSlot
}

// String return the name for debug.
func (dto GetEnclosureResponse) String() string {
	return dto.Name
}

// Load will load data from model.
func (dto *GetEnclosureResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.Enclosure)
	if !ok {
		log.Error("GetEnclosureResponsdto.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.DeviceIdentity = m.DeviceIdentity
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	// blade
	dto.ServerSlots = make([]ServerSlot, 0)
	for _, v := range m.ServerSlots {
		k := ServerSlot{}
		k.Load(&v)
		dto.ServerSlots = append(dto.ServerSlots, k)
	}
	// switch
	dto.SwitchSlots = make([]SwitchSlot, 0)
	for _, v := range m.SwitchSlots {
		k := SwitchSlot{}
		k.Load(&v)
		dto.SwitchSlots = append(dto.SwitchSlots, k)
	}
	// manager
	dto.ManagerSlots = make([]ManagerSlot, 0)
	for _, v := range m.ManagerSlots {
		k := ManagerSlot{}
		k.Load(&v)
		dto.ManagerSlots = append(dto.ManagerSlots, k)
	}
	// appliance
	dto.ApplianceSlots = make([]ApplianceSlot, 0)
	for _, v := range m.ApplianceSlots {
		k := ApplianceSlot{}
		k.Load(&v)
		dto.ApplianceSlots = append(dto.ApplianceSlots, k)
	}
	// fan
	dto.FanSlots = make([]FanSlot, 0)
	for _, v := range m.FanSlots {
		k := FanSlot{}
		k.Load(&v)
		dto.FanSlots = append(dto.FanSlots, k)
	}
	// power
	dto.PowerSlots = make([]PowerSlot, 0)
	for _, v := range m.PowerSlots {
		k := PowerSlot{}
		k.Load(&v)
		dto.PowerSlots = append(dto.PowerSlots, k)
	}
	return nil
}
