package model

// SupportedLinkCapabilities The self-described link capabilities of an assocaited port.
type SupportedLinkCapabilities struct {
	LinkNetworkTechnology string // The self-described link network technology capabilities of this port.
	LinkSpeedMbps         int    // The speed of the link in Mbps when this link network technology is active.
}

// NetDevFuncMaxBWAlloc A maximum bandwidth allocation percentage for a Network Device Functions associated a port.
type NetDevFuncMaxBWAlloc struct {
	MaxBWAllocPercent int // The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	// ???
	NetworkDeviceFunction NetworkDeviceFunction // Contains the members of this collection.
}

// NetDevFuncMinBWAlloc A minimum bandwidth allocation percentage for a Network Device Functions associated a port.
type NetDevFuncMinBWAlloc struct {
	MinBWAllocPercent int // The minimum bandwidth allocation percentage allocated to the corresponding network device function instance.
	// ???
	NetworkDeviceFunction NetworkDeviceFunction // Contains the members of this collection.
}

// NetworkPort A Network Port represents a discrete physical port capable of connecting to a network.
type NetworkPort struct {
	Resource
	PhysicalPortNumber         string   // The physical port number label for this port.
	LinkStatus                 string   // The status of the link between this port and its link partner.
	AssociatedNetworkAddresses []string // The array of configured network addresses (MAC or WWN) that are associated with this Network Port, including the programmed address of the lowest numbered Network Device Function, the configured but not active address if applicable, the address for hardware port teaming, or other network addresses.
}
