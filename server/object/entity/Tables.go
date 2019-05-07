package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Location", Info: new(Location)},
		{Name: "PostalAddress", Info: new(PostalAddress)},
		{Name: "Placement", Info: new(Placement)},
		{Name: "Server", Info: new(Server)},
		{Name: "Processor", Info: new(Processor)},
		{Name: "Memory", Info: new(Memory)},
		{Name: "EthernetInterface", Info: new(EthernetInterface)},
		{Name: "IPv4Address", Info: new(IPv4Address)},
		{Name: "IPv6Address", Info: new(IPv6Address)},
		{Name: "VLanNetworkInterface", Info: new(VLanNetworkInterface)},
		{Name: "NetworkInterface", Info: new(NetworkInterface)},
		{Name: "Storage", Info: new(Storage)},
		{Name: "StorageController", Info: new(StorageController)},
		{Name: "Power", Info: new(Power)},
		{Name: "Thermal", Info: new(Thermal)},
		{Name: "Fan", Info: new(Fan)},
		{Name: "PowerControl", Info: new(PowerControl)},
		{Name: "PowerSupply", Info: new(PowerSupply)},
		{Name: "Redundancy", Info: new(Redundancy)},
		{Name: "Board", Info: new(Board)},
		{Name: "NetworkAdapter", Info: new(NetworkAdapter)},
		{Name: "Controller", Info: new(Controller)},
		{Name: "Drive", Info: new(Drive)},
		{Name: "PCIeDevice", Info: new(PCIeDevice)},
		{Name: "PCIeFunction", Info: new(PCIeFunction)},
		{Name: "NetworkPort", Info: new(NetworkPort)},
		{Name: "ServerGroup", Info: new(ServerGroup)},
		{Name: "ServerServerGroup", Info: new(ServerServerGroup)},
		{Name: "AdapterConfig", Info: new(AdapterConfig)},
		{Name: "AdapterModel", Info: new(AdapterModel)},
		{Name: "AdapterCapability", Info: new(AdapterCapability)},
		{Name: "RAIDCapability", Info: new(RAIDCapability)},
		{Name: "EthernetCapability", Info: new(EthernetCapability)},
		{Name: "FCoECapability", Info: new(FCoECapability)},
	}

	// ServerTables The tables used by server resources.
	ServerTables = []base.TableInfo{
		{Name: "Location", Info: new(Location)},
		{Name: "PostalAddress", Info: new(PostalAddress)},
		{Name: "Placement", Info: new(Placement)},
		{Name: "Server", Info: new(Server)},
		{Name: "Processor", Info: new(Processor)},
		{Name: "Memory", Info: new(Memory)},
		{Name: "EthernetInterface", Info: new(EthernetInterface)},
		{Name: "IPv4Address", Info: new(IPv4Address)},
		{Name: "IPv6Address", Info: new(IPv6Address)},
		{Name: "VLanNetworkInterface", Info: new(VLanNetworkInterface)},
		{Name: "NetworkInterface", Info: new(NetworkInterface)},
		{Name: "Storage", Info: new(Storage)},
		{Name: "StorageController", Info: new(StorageController)},
		{Name: "Power", Info: new(Power)},
		{Name: "Thermal", Info: new(Thermal)},
		{Name: "Fan", Info: new(Fan)},
		{Name: "PowerControl", Info: new(PowerControl)},
		{Name: "PowerSupply", Info: new(PowerSupply)},
		{Name: "Redundancy", Info: new(Redundancy)},
		{Name: "Board", Info: new(Board)},
		{Name: "NetworkAdapter", Info: new(NetworkAdapter)},
		{Name: "Controller", Info: new(Controller)},
		{Name: "Drive", Info: new(Drive)},
		{Name: "PCIeDevice", Info: new(PCIeDevice)},
		{Name: "PCIeFunction", Info: new(PCIeFunction)},
		{Name: "NetworkPort", Info: new(NetworkPort)},
	}
)
