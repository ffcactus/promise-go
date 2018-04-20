package model

// ControllerCapabilities The capabilities of a controller.
type ControllerCapabilities struct {
	NetworkPortCount int // The number of physical ports on this controller.
}

// Controller A network controller ASIC that makes up part of a NetworkAdapter.
type Controller struct {
	FirmwarePackageVersion string                 // The version of the user-facing firmware package.
	ControllerCapabilities ControllerCapabilities // The capabilities of this controller.
	NetworkPorts           []NetworkPort          // Contains the members of this collection.

}

// NetworkAdapter A NetworkAdapter represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	Resource
	ProductInfo
	Controllers []Controller // The set of network controllers ASICs that make up this NetworkAdapter.
}
