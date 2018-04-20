package model

// NetworkInterface A NetworkInterface contains references linking NetworkAdapter, NetworkPort, and NetworkDeviceFunction resources and represents the functionality available to the containing system.
type NetworkInterface struct {
	Resource
	NetworkAdapterURI string
}
