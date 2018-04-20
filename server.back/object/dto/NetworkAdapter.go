package dto

import (
	"promise/server/object/model"
)

// ControllerCapabilities The capabilities of a controller.
type ControllerCapabilities struct {
	NetworkPortCount int `json:"NetworkPortCount"` // The number of physical ports on this controller.
	// NetworkDeviceFunctionCount int                   // The maximum number of physical functions available on this controller.
	// DataCenterBridging         DataCenterBridging    // Data Center Bridging (DCB) for this controller.
	// VirtualizationOffload      VirtualizationOffload // Virtualization offload for this controller.
	// NPIV                       NPIV                  // N_Port ID Virtualization (NPIV) capabilties for this controller.
}

// Controller A network controller ASIC that makes up part of a NetworkAdapter.
type Controller struct {
	FirmwarePackageVersion string                 `json:"FirmwarePackageVersion"` // The version of the user-facing firmware package.
	ControllerCapabilities ControllerCapabilities `json:"Slot"`                   // The capabilities of this controller.
	NetworkPorts           []NetworkPort          `json:"NetworkPorts"`           // Contains the members of this collection.
}

// Load will load data from model.
func (dto *Controller) Load(m *model.Controller) {
	dto.FirmwarePackageVersion = m.FirmwarePackageVersion
	dto.ControllerCapabilities.NetworkPortCount = m.ControllerCapabilities.NetworkPortCount
	for i := range m.NetworkPorts {
		portD := NetworkPort{}
		portD.Load(&m.NetworkPorts[i])
		dto.NetworkPorts = append(dto.NetworkPorts, portD)
	}
}

// NetworkAdapter represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	ResourceResponse
	ProductInfoResponse
	Controllers []Controller `json:"Controllers"` // The set of network controllers ASICs that make up this NetworkAdapter.
	//	NetworkPorts []NetworkPort // Contains the members of this collection.
	//	NetworkDeviceFunctions []NetworkDeviceFunction // Contains the members of this collection.
}

// Load will load data from model.
func (dto *NetworkAdapter) Load(m *model.NetworkAdapter) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	for i := range m.Controllers {
		controllerD := Controller{}
		controllerD.Load(&m.Controllers[i])
		dto.Controllers = append(dto.Controllers, controllerD)
	}
}
