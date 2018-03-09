package dto

import (
	"promise/server/object/model"
)

type ComputerSystem struct {
	Processors         []Processor         `json:"Processors"`
	Memory             []Memory            `json:"Memory"`
	EthernetInterfaces []EthernetInterface `json:"EthernetInterfaces"`
	NetworkInterfaces  []NetworkInterface  `json:"NetworkInterfaces"`
	Storages           []Storage           `json:"Storages"`
}

type Chassis struct {
	Power           Power            `json:"Power"`
	Thermal         Thermal          `json:"Thermal"`
	OemHuaweiBoards []OemHuaweiBoard `json:"OemHuaweiBoards"`
	NetworkAdapters []NetworkAdapter `json:"NetworkAdapters"`
	Drives          []Drive          `json:"Drives"`
	PCIeDevices     []PCIeDevice     `json:"PCIeDevices"`
}

type GetServerResponse struct {
	ID             string         `json:"ID"`
	URI            string         `json:"URI"`
	Name           string         `json:"Name"`
	Description    string         `json:"Description"`
	State          string         `json:"State"`
	Health         string         `json:"Health"`
	PhysicalUUID   string         `json:"PhysicalUUID"`
	Address        string         `json:"Address"`
	Type           string         `json:"Type"`
	CurrentTask    string         `json:"CurrentTask"`
	ComputerSystem ComputerSystem `json:"ComputerSystem"`
	Chassis        Chassis        `json:"Chassis"`
}

func (dto *GetServerResponse) Load(m *model.Server) {
	if m == nil {
		return
	}
	dto.ID = m.ID
	dto.URI = m.URI
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	dto.PhysicalUUID = m.PhysicalUUID
	dto.Address = m.Address
	dto.Type = m.Type
	dto.CurrentTask = m.CurrentTask
	// ComputerSystem.Processors
	dto.ComputerSystem.Processors = make([]Processor, 0)
	for i := range m.ComputerSystem.Processors {
		each := new(Processor)
		each.Load(&m.ComputerSystem.Processors[i])
		dto.ComputerSystem.Processors = append(dto.ComputerSystem.Processors, *each)
	}

	// ComputerSystem.Memory
	dto.ComputerSystem.Memory = make([]Memory, 0)
	for i := range m.ComputerSystem.Memory {
		each := new(Memory)
		each.Load(&m.ComputerSystem.Memory[i])
		dto.ComputerSystem.Memory = append(dto.ComputerSystem.Memory, *each)
	}

	// ComputerSystem.EthernetInterfaces
	dto.ComputerSystem.EthernetInterfaces = make([]EthernetInterface, 0)
	for i := range m.ComputerSystem.EthernetInterfaces {
		each := new(EthernetInterface)
		each.Load(&m.ComputerSystem.EthernetInterfaces[i])
		dto.ComputerSystem.EthernetInterfaces = append(dto.ComputerSystem.EthernetInterfaces, *each)
	}
	// ComputerSystem.NetworkInterfaces
	dto.ComputerSystem.NetworkInterfaces = make([]NetworkInterface, 0)
	for i := range m.ComputerSystem.NetworkInterfaces {
		each := new(NetworkInterface)
		each.Load(&m.ComputerSystem.NetworkInterfaces[i], m.Chassis.NetworkAdapters)
		dto.ComputerSystem.NetworkInterfaces = append(dto.ComputerSystem.NetworkInterfaces, *each)
	}
	// ComputerSystem.Storages
	dto.ComputerSystem.Storages = make([]Storage, 0)
	for i := range m.ComputerSystem.Storages {
		each := new(Storage)
		each.Load(&m.ComputerSystem.Storages[i], m.Chassis.Drives)
		dto.ComputerSystem.Storages = append(dto.ComputerSystem.Storages, *each)
	}
	// Chassis.Power
	dto.Chassis.Power.Load(&m.Chassis.Power)

	// Chassis.Thermal
	dto.Chassis.Thermal.Load(&m.Chassis.Thermal)

	// Chassis.OemHuaweiBoards
	dto.Chassis.OemHuaweiBoards = make([]OemHuaweiBoard, 0)
	for i := range m.Chassis.OemHuaweiBoards {
		each := new(OemHuaweiBoard)
		each.Load(&m.Chassis.OemHuaweiBoards[i])
		dto.Chassis.OemHuaweiBoards = append(dto.Chassis.OemHuaweiBoards, *each)
	}
	// Chassis.NetworkAdapters
	dto.Chassis.NetworkAdapters = make([]NetworkAdapter, 0)
	for i := range m.Chassis.NetworkAdapters {
		each := new(NetworkAdapter)
		each.Load(&m.Chassis.NetworkAdapters[i])
		dto.Chassis.NetworkAdapters = append(dto.Chassis.NetworkAdapters, *each)
	}

	// Chassis.Drives
	dto.Chassis.Drives = make([]Drive, 0)
	for i := range m.Chassis.Drives {
		each := new(Drive)
		each.Load(&m.Chassis.Drives[i])
		dto.Chassis.Drives = append(dto.Chassis.Drives, *each)
	}

	// Chassis.PCIeDevices
	dto.Chassis.PCIeDevices = make([]PCIeDevice, 0)
	for i := range m.Chassis.PCIeDevices {
		each := new(PCIeDevice)
		each.Load(&m.Chassis.PCIeDevices[i], m.ComputerSystem.EthernetInterfaces)
		dto.Chassis.PCIeDevices = append(dto.Chassis.PCIeDevices, *each)
	}
}
