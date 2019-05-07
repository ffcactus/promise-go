package dto

// A entry in the RFC 6724 Address Selection Policy Table.
type IPv6AddressPolicyEntry struct {
	Prefix     *string `json:"Prefix"`     // The IPv6 Address Prefix (as defined in RFC 6724 section 2.1).
	Precedence *int    `json:"Precedence"` // The IPv6 Precedence (as defined in RFC 6724 section 2.1.
	Label      *int    `json:"Label"`      // The IPv6 Label (as defined in RFC 6724 section 2.1).
}

// Contains references to other resources that are related to this resource.
type EthernetInterfaceLinks struct {
	Endpoints     *[]string `json:"Endpoints"`     // An array of references to the endpoints that connect to this ethernet interface.
	HostInterface *string   `json:"HostInterface"` // This is a reference to a Host Interface that is associated with this Ethernet Interface.
	Chassis       *string   `json:"Chassis"`       // A reference to the Chassis which contains this Ethernet Interface.
}

// This schema defines a simple ethernet NIC resource.
type GetEthernetInterfaceResponse struct {
	Resource
	Links                  *EthernetInterfaceLinks   `json:"Links"`                  // Contains references to other resources that are related to this resource.
	UefiDevicePath         *string                   `json:"UefiDevicePath"`         // The UEFI device path for this interface.
	InterfaceEnabled       *bool                     `json:"InterfaceEnabled"`       // This indicates whether this interface is enabled.
	PermanentMACAddress    *string                   `json:"PermanentMACAddress"`    // This is the permanent MAC address assigned to this interface (port).
	MACAddress             *string                   `json:"MACAddress"`             // This is the currently configured MAC address of the (logical port) interface.
	SpeedMbps              *int                      `json:"SpeedMbps"`              // This is the current speed in Mbps of this interface.
	AutoNeg                *bool                     `json:"AutoNeg"`                // This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	FullDuplex             *bool                     `json:"FullDuplex"`             // This indicates if the interface is in Full Duplex mode or not.
	MTUSize                *int                      `json:"MTUSize"`                // This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	HostName               *string                   `json:"HostName"`               // The DNS Host Name, without any domain information.
	FQDN                   *string                   `json:"FQDN"`                   // This is the complete, fully qualified domain name obtained by DNS for this interface.
	MaxIPv6StaticAddresses *string                   `json:"MaxIPv6StaticAddresses"` // This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.
	VLAN                   *VLAN                     `json:"VLAN"`                   // If this Network Interface supports more than one VLAN, this property will not be present and the client should look for VLANs collection in the link section of this resource.
	IPv4Addresses          *[]IPv4Address            `json:"IPv4Addresses"`          // The IPv4 addresses assigned to this interface.
	IPv6AddressPolicyTable *[]IPv6AddressPolicyEntry `json:"IPv6AddressPolicyTable"` // An array representing the RFC 6724 Address Selection Policy Table.
	IPv6Addresses          *[]IPv6Address            `json:"IPv6Addresses"`          // This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.
	IPv6StaticAddresses    *[]IPv6Address            `json:"IPv6StaticAddresses"`    // This array of objects represents all of the IPv6 static addresses to be assigned on this interface.
	IPv6DefaultGateway     *string                   `json:"IPv6DefaultGateway"`     // This is the IPv6 default gateway address that is currently in use on this interface.
	NameServers            *[]string                 `json:"NameServers"`            // This represents DNS name servers that are currently in use on this interface.
	VLANs                  *OdataID                  `json:"VLANs"`                  // This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.
	LinkStatus             *string                   `json:"LinkStatus"`             // The link status of this interface (port).
}
