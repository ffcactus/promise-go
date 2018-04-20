package entity

// NetworkPort A Network Port represents a discrete physical port capable of connecting to a network.
type NetworkPort struct {
	EmbeddedResource
	Ref                        uint
	PhysicalPortNumber         string // The physical port number label for this port.
	LinkStatus                 string // The status of the link between this port and its link partner.
	AssociatedNetworkAddresses string // The array of configured network addresses (MAC or WWN) that are associated with this Network Port, including the programmed address of the lowest numbered Network Device Function, the configured but not active address if applicable, the address for hardware port teaming, or other network addresses.
}
