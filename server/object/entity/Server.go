package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// Server Server entity.
type Server struct {
	base.Entity
	State              string              `gorm:"column:State"`
	Health             string              `gorm:"column:Health"`
	Name               string              `gorm:"column:Name"`
	Description        string              `gorm:"column:Description"`
	OriginURIsChassis  *string             `gorm:"column:OriginURIsChassis"`
	OriginURIsSystem   *string             `gorm:"column:OriginURIsSystem"`
	PhysicalUUID       string              `gorm:"column:PhysicalUUID"`
	Hostname           string              `gorm:"column:Hostname"`
	Type               string              `gorm:"column:Type"`
	Protocol           string              `gorm:"column:Protocol"`
	Credential         string              `gorm:"column:Credential"`
	Processors         []Processor         `gorm:"column:Processors;ForeignKey:ServerRef"`
	Memory             []Memory            `gorm:"column:Memory;ForeignKey:ServerRef"`
	EthernetInterfaces []EthernetInterface `gorm:"column:EthernetInterfaces;ForeignKey:ServerRef"`
	NetworkInterfaces  []NetworkInterface  `gorm:"column:NetworkInterfaces;ForeignKey:ServerRef"`
	Storages           []Storage           `gorm:"column:Storages;ForeignKey:ServerRef"`
	Power              Power               `gorm:"column:Power;ForeignKey:ServerRef"`
	Thermal            Thermal             `gorm:"column:Thermal;ForeignKey:ServerRef"`
	OemHuaweiBoards    []OemHuaweiBoard    `gorm:"column:OemHuaweiBoards;ForeignKey:ServerRef"`
	Drives             []Drive             `gorm:"column:Drives;ForeignKey:ServerRef"`
	PCIeDevices        []PCIeDevice        `gorm:"column:PCIeDevices;ForeignKey:ServerRef"`
	NetworkAdapters    []NetworkAdapter    `gorm:"column:NetworkAdapters;ForeignKey:ServerRef"`
}

// TableName will set the table name.
func (Server) TableName() string {
	return "Server"
}

// DebugInfo return the debug name of this entity.
func (e *Server) DebugInfo() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Server) PropertyNameForDuplicationCheck() string {
	return "PhysicalUUID"
}

// Preload return the property names that need to be preload.
func (e *Server) Preload() []string {
	return []string{
		"Processors",
		"Memory",
		"EthernetInterfaces",
		"EthernetInterfaces.IPv4Addresses",
		"EthernetInterfaces.IPv6Addresses",
		"EthernetInterfaces.VLANs",
		"NetworkInterfaces",
		"Storages",
		"Storages.StorageControllers",
		"Power",
		"Power.PowerControl",
		"Power.Voltages",
		"Power.PowerSupplies",
		"Power.Redundancy",
		"Thermal",
		"Thermal.Temperatures",
		"Thermal.Fans",
		"OemHuaweiBoards",
		"NetworkAdapters",
		"NetworkAdapters.Controllers",
		"NetworkAdapters.Controllers.NetworkPorts",
		"Drives",
		"Drives.Location",
		"Drives.Location.PostalAddress",
		"Drives.Location.Placement",
		"PCIeDevices",
		"PCIeDevices.PCIeFunctions",
	}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *Server) Association() []interface{} {
	ret := []interface{}{}
	for _, x := range e.Processors {
		ret = append(ret, &x)
	}
	for _, x := range e.Memory {
		ret = append(ret, &x)
	}
	for _, x := range e.EthernetInterfaces {
		for _, y := range x.IPv4Addresses {
			ret = append(ret, &y)
		}
		for _, y := range x.IPv6Addresses {
			ret = append(ret, &y)
		}
		for _, y := range x.VLANs {
			ret = append(ret, &y)
		}
		ret = append(ret, &x)
	}
	for _, x := range e.NetworkInterfaces {
		ret = append(ret, &x)
	}
	for _, x := range e.Storages {
		for _, y := range x.StorageControllers {
			ret = append(ret, &y)
		}
		ret = append(ret, &x)
	}
	for _, x := range e.Power.PowerControl {
		ret = append(ret, &x)
	}
	for _, x := range e.Power.Voltages {
		ret = append(ret, &x)
	}
	for _, x := range e.Power.PowerSupplies {
		ret = append(ret, &x)
	}
	for _, x := range e.Power.Redundancy {
		ret = append(ret, &x)
	}
	ret = append(ret, &e.Power)
	for _, x := range e.Thermal.Temperatures {
		ret = append(ret, &x)
	}
	for _, x := range e.Thermal.Fans {
		ret = append(ret, &x)
	}
	ret = append(ret, &e.Thermal)
	for _, x := range e.OemHuaweiBoards {
		ret = append(ret, &x)
	}
	for _, x := range e.NetworkAdapters {
		for _, y := range x.Controllers {
			for _, z := range y.NetworkPorts {
				ret = append(ret, &z)
			}
			ret = append(ret, &y)
		}
		ret = append(ret, &x)
	}
	for _, x := range e.Drives {
		for _, y := range x.Location {
			ret = append(ret, y.PostalAddress)
			ret = append(ret, y.Placement)
			ret = append(ret, &y)
		}
		ret = append(ret, &x)
	}
	for _, x := range e.PCIeDevices {
		for _, y := range x.PCIeFunctions {
			ret = append(ret, &y)
		}
		ret = append(ret, &x)
	}
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *Server) Tables() []interface{} {
	return []interface{}{
		new(Location),
		new(PostalAddress),
		new(Placement),
		new(Server),
		new(Processor),
		new(Memory),
		new(EthernetInterface),
		new(IPv4Address),
		new(IPv6Address),
		new(VLanNetworkInterface),
		new(NetworkInterface),
		new(Storage),
		new(StorageController),
		new(Power),
		new(Thermal),
		new(Temperature),
		new(Fan),
		new(PowerControl),
		new(Voltage),
		new(PowerSupply),
		new(Redundancy),
		new(OemHuaweiBoard),
		new(NetworkAdapter),
		new(Controller),
		new(Drive),
		new(PCIeDevice),
		new(PCIeFunction),
		new(NetworkPort),
	}
}

// FilterNameList return all the property name that can be used in filter.
func (e *Server) FilterNameList() []string {
	return []string{
		"State",
		"Name",
		"Health",
		"Description",
		"PhysicalUUID",
		"Hostname",
	}
}

// Load will load data from model. this function is used on POST.
func (e *Server) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Server)
	if !ok {
		log.Error("entity.Server.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
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
	return nil
}

// ToModel will create a new model from entity.
func (e *Server) ToModel() base.ModelInterface {
	m := model.Server{}
	base.EntityToModel(&e.Entity, &m.Model)
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
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *Server) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.ServerCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	m.State = e.State
	m.Health = e.Health
	return m
}
