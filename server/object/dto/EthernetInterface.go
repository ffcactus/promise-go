package dto

import (
	"promise/server/object/model"
)

// This schema defines a simple ethernet NIC resource.
type EthernetInterface struct {
	ResourceResponse
	UefiDevicePath         *string                `json:"UefiDevicePath,omitempty"`         // The UEFI device path for this interface.
	InterfaceEnabled       *bool                  `json:"InterfaceEnabled,omitempty"`       // This indicates whether this interface is enabled.
	PermanentMACAddress    *string                `json:"PermanentMACAddress,omitempty"`    // This is the permanent MAC address assigned to this interface (port).
	MACAddress             *string                `json:"MACAddress,omitempty"`             // This is the currently configured MAC address of the (logical port) interface.
	SpeedMbps              *int                   `json:"SpeedMbps,omitempty"`              // This is the current speed in Mbps of this interface.
	AutoNeg                *bool                  `json:"AutoNeg,omitempty"`                // This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	FullDuplex             *bool                  `json:"FullDuplex,omitempty"`             // This indicates if the interface is in Full Duplex mode or not.
	MTUSize                *int                   `json:"MTUSize,omitempty"`                // This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	HostName               *string                `json:"HostName,omitempty"`               // The DNS Host Name, without any domain information.
	FQDN                   *string                `json:"FQDN,omitempty"`                   // This is the complete, fully qualified domain name obtained by DNS for this interface.
	MaxIPv6StaticAddresses *string                `json:"MaxIPv6StaticAddresses,omitempty"` // This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.
	IPv4Addresses          []IPv4Address          `json:"IPv4Addresses"`                    // The IPv4 addresses assigned to this interface.
	IPv6Addresses          []IPv6Address          `json:"IPv6Addresses"`                    // This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.
	IPv6DefaultGateway     *string                `json:"IPv6DefaultGateway,omitempty"`     // This is the IPv6 default gateway address that is currently in use on this interface.
	VLANs                  []VLanNetworkInterface `json:"VLANs"`                            // This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.
	LinkStatus             *string                `json:"LinkStatus,omitempty"`             // The link status of this interface (port).
}

// This type describes an IPv4 Address.
type IPv4Address struct {
	Address       *string `json:"Address,omitempty"`       // This is the IPv4 Address.
	SubnetMask    *string `json:"SubnetMask,omitempty"`    // This is the IPv4 Subnet mask.
	AddressOrigin *string `json:"AddressOrigin,omitempty"` // This indicates how the address was determined.
	Gateway       *string `json:"Gateway,omitempty"`       // This is the IPv4 gateway for this address.
}

// This type describes an IPv6 Address.
type IPv6Address struct {
	Address       *string `json:"Address,omitempty"`       // This is the IPv6 Address.
	PrefixLength  *int    `json:"PrefixLength,omitempty"`  // This is the IPv6 Address Prefix Length.
	AddressOrigin *string `json:"AddressOrigin,omitempty"` // This indicates how the address was determined.
	AddressState  *string `json:"AddressState,omitempty"`  // The current state of this address as defined in RFC 4862.
}

type VLanNetworkInterface struct {
	VLANEnable *bool `json:"VLANEnable,omitempty"` // This indicates if this VLAN is enabled.
	VLANID     *int  `json:"VLANID,omitempty"`     // This indicates the VLAN identifier for this VLAN.
}

func (this *EthernetInterface) Load(m *model.EthernetInterface) {
	this.LoadResourceResponse(&(*m).Resource)
	this.UefiDevicePath = m.UefiDevicePath
	this.InterfaceEnabled = m.InterfaceEnabled
	this.PermanentMACAddress = m.PermanentMACAddress
	this.MACAddress = m.MACAddress
	this.SpeedMbps = m.SpeedMbps
	this.AutoNeg = m.AutoNeg
	this.FullDuplex = m.FullDuplex
	this.MTUSize = m.MTUSize
	this.HostName = m.HostName
	this.FQDN = m.FQDN
	this.MaxIPv6StaticAddresses = m.MaxIPv6StaticAddresses
	this.LinkStatus = m.LinkStatus
	this.IPv4Addresses = make([]IPv4Address, 0)
	if m.IPv4Addresses != nil {
		for i, _ := range m.IPv4Addresses {
			each := IPv4Address{}
			each.Address = m.IPv4Addresses[i].Address
			each.SubnetMask = m.IPv4Addresses[i].SubnetMask
			each.AddressOrigin = m.IPv4Addresses[i].AddressOrigin
			each.Gateway = m.IPv4Addresses[i].Gateway
			this.IPv4Addresses = append(this.IPv4Addresses, each)
		}
	}
	this.IPv6Addresses = make([]IPv6Address, 0)
	if m.IPv6Addresses != nil {
		for i, _ := range m.IPv6Addresses {
			each := IPv6Address{}
			each.Address = m.IPv6Addresses[i].Address
			each.PrefixLength = m.IPv6Addresses[i].PrefixLength
			each.AddressOrigin = m.IPv6Addresses[i].AddressOrigin
			each.AddressState = m.IPv6Addresses[i].AddressState
			this.IPv6Addresses = append(this.IPv6Addresses, each)
		}
	}
	this.VLANs = make([]VLanNetworkInterface, 0)
	if m.VLANs != nil {
		for i, _ := range m.VLANs {
			each := VLanNetworkInterface{}
			each.VLANEnable = m.VLANs[i].VLANEnable
			each.VLANID = m.VLANs[i].VLANID
			this.VLANs = append(this.VLANs, each)
		}
	}
}
