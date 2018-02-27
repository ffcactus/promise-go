package entity

// Controller A network controller ASIC that makes up part of a NetworkAdapter.
type Controller struct {
	NetworkAdapterRef uint
	EmbeddedObject
	FirmwarePackageVersion                 string        // The version of the user-facing firmware package.
	ControllerCapabilitiesNetworkPortCount int           // The capabilities of this controller.
	NetworkPorts                           []NetworkPort `gorm:"ForeignKey:Ref"`
}

// NetworkAdapter A NetworkAdapter represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	Controllers []Controller `gorm:"ForeignKey:NetworkAdapterRef"` // The set of network controllers ASICs that make up this NetworkAdapter.
}
