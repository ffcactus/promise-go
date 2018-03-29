package entity

import (
	"promise/common/object/entity"
	"promise/server/object/model"
)

// Server Server entity.
type Server struct {
	entity.PromiseEntity
	State              string
	Health             string
	Name               string
	Description        string
	OriginURIsChassis  *string
	OriginURIsSystem   *string
	PhysicalUUID       string
	Hostname           string
	Type               string
	Protocol           string
	Credential         string
	Processors         []Processor         `gorm:"ForeignKey:ServerRef"`
	Memory             []Memory            `gorm:"ForeignKey:ServerRef"`
	EthernetInterfaces []EthernetInterface `gorm:"ForeignKey:ServerRef"`
	NetworkInterfaces  []NetworkInterface  `gorm:"ForeignKey:ServerRef"`
	Storages           []Storage           `gorm:"ForeignKey:ServerRef"`
	Power              Power               `gorm:"ForeignKey:ServerRef"`
	Thermal            Thermal             `gorm:"ForeignKey:ServerRef"`
	OemHuaweiBoards    []OemHuaweiBoard    `gorm:"ForeignKey:ServerRef"`
	Drives             []Drive             `gorm:"ForeignKey:ServerRef"`
	PCIeDevices        []PCIeDevice        `gorm:"ForeignKey:ServerRef"`
	NetworkAdapters    []NetworkAdapter    `gorm:"ForeignKey:ServerRef"`
}

// ToModel will create a new model from entity.
func (e *Server) ToModel() *model.Server {
	m := new(model.Server)
	m.PromiseModel = e.PromiseEntity.ToModel()
	m.ID = e.ID
	m.OriginURIs.Chassis = e.OriginURIsChassis
	m.OriginURIs.System = e.OriginURIsSystem
	m.PhysicalUUID = e.PhysicalUUID
	m.Name = e.Name
	m.Description = e.Description
	m.Hostname = e.Hostname
	m.Type = e.Type
	m.Protocol = e.Protocol
	m.Credential = e.Credential
	m.State = e.State
	m.Health = e.Health
	// ComputerSystem.Processors
	processors := []model.Processor{}
	for i := range e.Processors {
		processors = append(processors, *e.Processors[i].ToModel())
	}
	m.ComputerSystem.Processors = processors
	// ComputerSystem.Memory
	memory := []model.Memory{}
	for i := range e.Memory {
		memory = append(memory, *e.Memory[i].ToModel())
	}
	m.ComputerSystem.Memory = memory

	// ComputerSystem.EthernetInterfaces
	ethernetInterfaces := []model.EthernetInterface{}
	for i := range e.EthernetInterfaces {
		ethernetInterfaces = append(ethernetInterfaces, *e.EthernetInterfaces[i].ToModel())
	}
	m.ComputerSystem.EthernetInterfaces = ethernetInterfaces
	// ComputerSystem.NetworkInterfaces
	networkInterfaces := []model.NetworkInterface{}
	for i := range e.NetworkInterfaces {
		networkInterfaces = append(networkInterfaces, *e.NetworkInterfaces[i].ToModel())
	}
	m.ComputerSystem.NetworkInterfaces = networkInterfaces
	// ComputerSystem.Storages
	storages := []model.Storage{}
	for i := range e.Storages {
		storages = append(storages, *e.Storages[i].ToModel())
	}
	m.ComputerSystem.Storages = storages
	// Chassis.Power
	createResourceModel(&e.Power.EmbeddedResource, &m.Chassis.Power.Resource)
	powerControl := []model.PowerControl{}
	for i := range e.Power.PowerControl {
		powerControl = append(powerControl, *e.Power.PowerControl[i].ToModel())
	}
	m.Chassis.Power.PowerControl = &powerControl

	voltages := []model.Voltage{}
	for i := range e.Power.Voltages {
		voltages = append(voltages, *e.Power.Voltages[i].ToModel())
	}
	m.Chassis.Power.Voltages = &voltages

	powerSupplies := []model.PowerSupply{}
	for i := range e.Power.PowerSupplies {
		powerSupplies = append(powerSupplies, *e.Power.PowerSupplies[i].ToModel())
	}
	m.Chassis.Power.PowerSupplies = &powerSupplies

	redundancy := []model.Redundancy{}
	for i := range e.Power.Redundancy {
		redundancy = append(redundancy, *e.Power.Redundancy[i].ToModel())
	}
	m.Chassis.Power.Redundancy = &redundancy
	// Chassis.Thermal
	createResourceModel(&e.Thermal.EmbeddedResource, &m.Chassis.Thermal.Resource)
	temperatures := []model.Temperature{}
	for i := range e.Thermal.Temperatures {
		temperatures = append(temperatures, *e.Thermal.Temperatures[i].ToModel())
	}
	m.Chassis.Thermal.Temperatures = temperatures
	fans := []model.Fan{}
	for i := range e.Thermal.Fans {
		fans = append(fans, *e.Thermal.Fans[i].ToModel())
	}
	m.Chassis.Thermal.Fans = fans
	// Chassis.OemHuaweiBoards
	boards := []model.OemHuaweiBoard{}
	for i := range e.OemHuaweiBoards {
		boards = append(boards, *e.OemHuaweiBoards[i].ToModel())
	}
	m.Chassis.OemHuaweiBoards = boards
	// Chassis.NetworkAdapters
	networkAdapters := []model.NetworkAdapter{}
	for i := range e.NetworkAdapters {
		networkAdapters = append(networkAdapters, *e.NetworkAdapters[i].ToModel())
	}
	m.Chassis.NetworkAdapters = networkAdapters
	// Chassis.Drives
	drives := []model.Drive{}
	for i := range e.Drives {
		drives = append(drives, *e.Drives[i].ToModel())
	}
	m.Chassis.Drives = drives
	// Chassis.PCIeDevices
	pcieDevices := []model.PCIeDevice{}
	for i := range e.PCIeDevices {
		pcieDevices = append(pcieDevices, *e.PCIeDevices[i].ToModel())
	}
	m.Chassis.PCIeDevices = pcieDevices
	return m
}

// Load will load data from model.
func (e *Server) Load(m *model.Server) {
	e.PromiseEntity.Load(m.PromiseModel)
	e.State = m.State
	e.Health = m.Health
	e.Name = m.Name
	e.Description = m.Description
	e.OriginURIsChassis = m.OriginURIs.Chassis
	e.OriginURIsSystem = m.OriginURIs.System
	e.PhysicalUUID = m.PhysicalUUID
	e.Hostname = m.Hostname
	e.Credential = m.Credential
	e.Type = m.Type
	e.Protocol = m.Protocol
}
