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

func (this *GetServerResponse) Load(m *model.Server) {
	if m == nil {
		return
	}
	this.ID = m.ID
	this.URI = m.URI
	this.Name = m.Name
	this.Description = m.Description
	this.State = m.State
	this.Health = m.Health
	this.PhysicalUUID = m.PhysicalUUID
	this.Address = m.Address
	this.Type = m.Type
	this.CurrentTask = m.CurrentTask
	// ComputerSystem.Processors
	this.ComputerSystem.Processors = make([]Processor, 0)
	for i := range m.ComputerSystem.Processors {
		each := new(Processor)
		each.Load(&m.ComputerSystem.Processors[i])
		this.ComputerSystem.Processors = append(this.ComputerSystem.Processors, *each)
	}

	// ComputerSystem.Memory
	this.ComputerSystem.Memory = make([]Memory, 0)
	for i := range m.ComputerSystem.Memory {
		each := new(Memory)
		each.Load(&m.ComputerSystem.Memory[i])
		this.ComputerSystem.Memory = append(this.ComputerSystem.Memory, *each)
	}

	// ComputerSystem.EthernetInterfaces
	this.ComputerSystem.EthernetInterfaces = make([]EthernetInterface, 0)
	for i := range m.ComputerSystem.EthernetInterfaces {
		each := new(EthernetInterface)
		each.Load(&m.ComputerSystem.EthernetInterfaces[i])
		this.ComputerSystem.EthernetInterfaces = append(this.ComputerSystem.EthernetInterfaces, *each)
	}
	// ComputerSystem.NetworkInterfaces
	this.ComputerSystem.NetworkInterfaces = make([]NetworkInterface, 0)
	for i := range m.ComputerSystem.NetworkInterfaces {
		each := new(NetworkInterface)
		each.Load(&m.ComputerSystem.NetworkInterfaces[i], m.Chassis.NetworkAdapters)
		this.ComputerSystem.NetworkInterfaces = append(this.ComputerSystem.NetworkInterfaces, *each)
	}
	// ComputerSystem.Storages
	this.ComputerSystem.Storages = make([]Storage, 0)
	for i := range m.ComputerSystem.Storages {
		each := new(Storage)
		each.Load(&m.ComputerSystem.Storages[i], m.Chassis.Drives)
		this.ComputerSystem.Storages = append(this.ComputerSystem.Storages, *each)
	}
	// Chassis.Power
	this.Chassis.Power.Load(&m.Chassis.Power)

	// Chassis.Thermal
	this.Chassis.Thermal.Load(&m.Chassis.Thermal)

	// Chassis.OemHuaweiBoards
	this.Chassis.OemHuaweiBoards = make([]OemHuaweiBoard, 0)
	for i := range m.Chassis.OemHuaweiBoards {
		each := new(OemHuaweiBoard)
		each.Load(&m.Chassis.OemHuaweiBoards[i])
		this.Chassis.OemHuaweiBoards = append(this.Chassis.OemHuaweiBoards, *each)
	}
	// Chassis.NetworkAdapters
	this.Chassis.NetworkAdapters = make([]NetworkAdapter, 0)
	for i := range m.Chassis.NetworkAdapters {
		each := new(NetworkAdapter)
		each.Load(&m.Chassis.NetworkAdapters[i])
		this.Chassis.NetworkAdapters = append(this.Chassis.NetworkAdapters, *each)
	}

	// Chassis.Drives
	this.Chassis.Drives = make([]Drive, 0)
	for i := range m.Chassis.Drives {
		each := new(Drive)
		each.Load(&m.Chassis.Drives[i])
		this.Chassis.Drives = append(this.Chassis.Drives, *each)
	}

	// Chassis.PCIeDevices
	this.Chassis.PCIeDevices = make([]PCIeDevice, 0)
	for i := range m.Chassis.PCIeDevices {
		each := new(PCIeDevice)
		each.Load(&m.Chassis.PCIeDevices[i], m.ComputerSystem.EthernetInterfaces)
		this.Chassis.PCIeDevices = append(this.Chassis.PCIeDevices, *each)
	}
}
