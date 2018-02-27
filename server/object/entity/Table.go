package entity

import (
	commonDB "promise/common/db"
)

var (
	// Tables The tables used by this project.
	Tables = []commonDB.TableInfo{
		{"Location", new(Location)},
		{"PostalAddress", new(PostalAddress)},
		{"Placement", new(Placement)},
		{"Server", new(Server)},
		{"Processor", new(Processor)},
		{"Memory", new(Memory)},
		{"EthernetInterface", new(EthernetInterface)},
		{"IPv4Address", new(IPv4Address)},
		{"IPv6Address", new(IPv6Address)},
		{"VLanNetworkInterface", new(VLanNetworkInterface)},
		{"NetworkInterface", new(NetworkInterface)},
		{"Storage", new(Storage)},
		{"StorageController", new(StorageController)},
		{"Power", new(Power)},
		{"Thermal", new(Thermal)},
		{"Temperature", new(Temperature)},
		{"Fan", new(Fan)},
		{"PowerControl", new(PowerControl)},
		{"Voltage", new(Voltage)},
		{"PowerSupply", new(PowerSupply)},
		{"Redundancy", new(Redundancy)},
		{"OemHuaweiBoard", new(OemHuaweiBoard)},
		{"NetworkAdapter", new(NetworkAdapter)},
		{"Controller", new(Controller)},
		{"Drive", new(Drive)},
		{"PCIeDevice", new(PCIeDevice)},
		{"PCIeFunction", new(PCIeFunction)},
		{"NetworkPort", new(NetworkPort)},
	}
)
