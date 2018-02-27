package model

// EthernetInterface This schema defines a simple ethernet NIC resource.
type EthernetInterface struct {
	Resource
	UefiDevicePath         *string                // The UEFI device path for this interface.
	InterfaceEnabled       *bool                  // This indicates whether this interface is enabled.
	PermanentMACAddress    *string                // This is the permanent MAC address assigned to this interface (port).
	MACAddress             *string                // This is the currently configured MAC address of the (logical port) interface.
	SpeedMbps              *int                   // This is the current speed in Mbps of this interface.
	AutoNeg                *bool                  // This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	FullDuplex             *bool                  // This indicates if the interface is in Full Duplex mode or not.
	MTUSize                *int                   // This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	HostName               *string                // The DNS Host Name, without any domain information.
	FQDN                   *string                // This is the complete, fully qualified domain name obtained by DNS for this interface.
	MaxIPv6StaticAddresses *string                // This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.
	IPv4Addresses          []IPv4Address          // The IPv4 addresses assigned to this interface.
	IPv6Addresses          []IPv6Address          // This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.
	VLANs                  []VLanNetworkInterface // This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.
	LinkStatus             *string                // The link status of this interface (port).
}
