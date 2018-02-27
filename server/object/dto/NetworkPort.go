package dto

import (
	"promise/server/object/model"
)

// The self-described link capabilities of an assocaited port.
type SupportedLinkCapabilities struct {
	LinkNetworkTechnology string // The self-described link network technology capabilities of this port.
	LinkSpeedMbps         int    // The speed of the link in Mbps when this link network technology is active.
}

// A maximum bandwidth allocation percentage for a Network Device Functions associated a port.
type NetDevFuncMaxBWAlloc struct {
	MaxBWAllocPercent int // The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	// ???
	// NetworkDeviceFunction NetworkDeviceFunction // Contains the members of this collection.
}

// A minimum bandwidth allocation percentage for a Network Device Functions associated a port.
type NetDevFuncMinBWAlloc struct {
	MinBWAllocPercent int // The minimum bandwidth allocation percentage allocated to the corresponding network device function instance.
	// ???
	// NetworkDeviceFunction NetworkDeviceFunction // Contains the members of this collection.
}

// A Network Port represents a discrete physical port capable of connecting to a network.
type NetworkPort struct {
	ResourceResponse
	PhysicalPortNumber         string   // The physical port number label for this port.
	LinkStatus                 string   // The status of the link between this port and its link partner.
	AssociatedNetworkAddresses []string // The array of configured network addresses (MAC or WWN) that are associated with this Network Port, including the programmed address of the lowest numbered Network Device Function, the configured but not active address if applicable, the address for hardware port teaming, or other network addresses.
	// SupportedLinkCapabilities     SupportedLinkCapabilities // The self-described link capabilities of this port.
	// ActiveLinkTechnology          string                    // Network Port Active Link Technology.
	// SupportedEthernetCapabilities string                    // The set of Ethernet capabilities that this port supports.
	// NetDevFuncMinBWAlloc          []NetDevFuncMinBWAlloc    // The array of minimum bandwidth allocation percentages for the Network Device Functions associated with this port.
	// NetDevFuncMaxBWAlloc          []NetDevFuncMaxBWAlloc    // The array of maximum bandwidth allocation percentages for the Network Device Functions associated with this port.
	// EEEEnabled                    bool                      // Whether IEEE 802.3az Energy Efficient Ethernet (EEE) is enabled for this network port.
	// WakeOnLANEnabled              bool                      // Whether Wake on LAN (WoL) is enabled for this network port.
	// PortMaximumMTU                int                       // The largest maximum transmission unit (MTU) that can be configured for this network port.
	// FlowControlStatus             string                    // The 802.3x flow control behavior negotiated with the link partner for this network port (Ethernet-only).
	// FlowControlConfiguration      string                    // The locally configured 802.3x flow control setting for this network port.
	// SignalDetected                bool                      // Whether or not the port has detected enough signal on enough lanes to establish link.
}

func (this *NetworkPort) Load(m *model.NetworkPort) {
	this.LoadResourceResponse(&m.Resource)
	this.PhysicalPortNumber = m.PhysicalPortNumber
	this.LinkStatus = m.LinkStatus
	this.AssociatedNetworkAddresses = m.AssociatedNetworkAddresses
}
