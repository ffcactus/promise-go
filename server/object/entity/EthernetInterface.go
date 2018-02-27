package entity

// IPv6AddressPolicyEntry A entry in the RFC 6724 Address Selection Policy Table.
type IPv6AddressPolicyEntry struct {
	Prefix     *string // The IPv6 Address Prefix (as defined in RFC 6724 section 2.1).
	Precedence *int    // The IPv6 Precedence (as defined in RFC 6724 section 2.1.
	Label      *int    // The IPv6 Label (as defined in RFC 6724 section 2.1).
}

// EthernetInterfaceLinks Contains references to other resources that are related to this resource.
type EthernetInterfaceLinks struct {
	Endpoints     *[]string // An array of references to the endpoints that connect to this ethernet interface.
	HostInterface *string   // This is a reference to a Host Interface that is associated with this Ethernet Interface.
	Chassis       *string   // A reference to the Chassis which contains this Ethernet Interface.
}

// EthernetInterface This schema defines a simple ethernet NIC resource.
type EthernetInterface struct {
	ServerRef string
	EmbeddedResource
	UefiDevicePath         *string       // The UEFI device path for this interface.
	InterfaceEnabled       *bool         // This indicates whether this interface is enabled.
	PermanentMACAddress    *string       // This is the permanent MAC address assigned to this interface (port).
	MACAddress             *string       // This is the currently configured MAC address of the (logical port) interface.
	SpeedMbps              *int          // This is the current speed in Mbps of this interface.
	AutoNeg                *bool         // This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	FullDuplex             *bool         // This indicates if the interface is in Full Duplex mode or not.
	MTUSize                *int          // This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	HostName               *string       // The DNS Host Name, without any domain information.
	FQDN                   *string       // This is the complete, fully qualified domain name obtained by DNS for this interface.
	MaxIPv6StaticAddresses *string       // This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.
	IPv4Addresses          []IPv4Address `gorm:"ForeignKey:EthernetInterfaceRef"` // The IPv4 addresses assigned to this interface.
	IPv6Addresses          []IPv6Address `gorm:"ForeignKey:EthernetInterfaceRef"` // This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.
	// IPv6StaticAddresses    []IPv6Address          `gorm:"ForeignKey:EthernetInterfaceRef"` // This array of objects represents all of the IPv6 static addresses to be assigned on this interface.
	IPv6DefaultGateway *string                // This is the IPv6 default gateway address that is currently in use on this interface.
	VLANs              []VLanNetworkInterface `gorm:"ForeignKey:EthernetInterfaceRef"` // This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.
	LinkStatus         *string                // The link status of this interface (port).
}

// IPv4Address This type describes an IPv4 Address.
type IPv4Address struct {
	EthernetInterfaceRef uint
	EmbeddedObject
	Address       *string // This is the IPv4 Address.
	SubnetMask    *string // This is the IPv4 Subnet mask.
	AddressOrigin *string // This indicates how the address was determined.
	Gateway       *string // This is the IPv4 gateway for this address.
}

// IPv6Address This type describes an IPv6 Address.
type IPv6Address struct {
	EthernetInterfaceRef uint
	EmbeddedObject
	Address       *string // This is the IPv6 Address.
	PrefixLength  *int    // This is the IPv6 Address Prefix Length.
	AddressOrigin *string // This indicates how the address was determined.
	AddressState  *string // The current state of this address as defined in RFC 4862.
}

// VLanNetworkInterface VLan network interface object.
type VLanNetworkInterface struct {
	EthernetInterfaceRef uint
	EmbeddedResource
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}
