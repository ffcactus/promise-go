package dto

type NetworkInterfaceLinks struct {
	NetworkAdapter ResourceRef //
}

// A NetworkInterface contains references linking NetworkAdapter, NetworkPort, and NetworkDeviceFunction resources and represents the functionality available to the containing system.
type GetNetworkInterfaceResponse struct {
	Resource
	Links        NetworkInterfaceLinks
	NetworkPorts ResourceRef
}
