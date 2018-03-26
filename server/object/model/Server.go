package model

import (
	"promise/common/object/model"
)

// OriginURIs The
type OriginURIs struct {
	Chassis *string
	System  *string
}

// ComputerSystem Computer system object.
type ComputerSystem struct {
	Processors         []Processor
	Memory             []Memory
	EthernetInterfaces []EthernetInterface
	NetworkInterfaces  []NetworkInterface
	Storages           []Storage
}

// Chassis Chassis object.
type Chassis struct {
	Power           Power
	Thermal         Thermal
	OemHuaweiBoards []OemHuaweiBoard
	NetworkAdapters []NetworkAdapter
	Drives          []Drive
	PCIeDevices     []PCIeDevice
}

// ServerBasicInfo It represents the basic info about a server(Rack, Enclosure, Blade, Switch)
type ServerBasicInfo struct {
	OriginURIs     OriginURIs // The URIs that we retrieve info from.
	PhysicalUUID   string
	Name           string
	Description    string
	Hostname       string
	Type           string
	Protocol       string
	OriginUsername *string
	OriginPassword *string
}

// Server Server object.
type Server struct {
	model.PromiseModel
	Name           string
	Description    string
	State          string
	Health         string
	OriginURIs     OriginURIs // The URIs that we retrieve info from.
	PhysicalUUID   string
	Hostname       string
	Type           string
	Protocol       string
	OriginUsername *string
	OriginPassword *string
	Credential     string
	CurrentTask    string
	ComputerSystem ComputerSystem
	Chassis        Chassis
}

// CreateServer Create servr object.
func (o *ServerBasicInfo) CreateServer() *Server {
	server := new(Server)
	server.Name = o.Name
	server.Description = o.Description
	server.State = "State???"
	server.Health = "Health???"
	server.OriginURIs.Chassis = o.OriginURIs.Chassis
	server.OriginURIs.System = o.OriginURIs.System
	server.PhysicalUUID = o.PhysicalUUID
	server.Hostname = o.Hostname
	server.Type = o.Type
	server.Protocol = o.Protocol
	server.OriginUsername = o.OriginUsername
	server.OriginPassword = o.OriginPassword
	return server
}
