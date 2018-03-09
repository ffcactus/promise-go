package dto

import (
	"promise/server/object/model"
)

// The links used in Controllers.
// type ControllersLinks struct {
// 	//	PCIeDevices            []string
// 	NetworkPorts []string
// 	//	NetworkDeviceFunctions []string
// }

// Data Center Bridging (DCB) for capabilities of a controller.
// type DataCenterBridging struct {
// 	Capable bool // Whether this controller is capable of Data Center Bridging (DCB).
// }

// A virtual function of a controller.
// type VirtualFunction struct {
// 	DeviceMaxCount         int // The maximum number of Virtual Functions (VFs) supported by this controller.
// 	NetworkPortMaxCount    int // The maximum number of Virtual Functions (VFs) supported per network port for this controller.
// 	MinAssignmentGroupSize int // The minimum number of Virtual Functions (VFs) that can be allocated or moved between physical functions for this controller.
// }

// Single-Root Input/Output Virtualization (SR-IOV) capabilities.
// type SRIOV struct {
// 	SRIOVVEPACapable bool // Whether this controller supports Single Root Input/Output Virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
// }

// N_Port ID Virtualization (NPIV) capabilties for a controller.
// type NPIV struct {
// 	MaxDeviceLogins int // The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
// 	MaxPortLogins   int // The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
// }

// A Virtualization offload capability of a controller.
// type VirtualizationOffload struct {
// 	VirtualFunction VirtualFunction // A virtual function of a controller.
// 	SRIOV           SRIOV           // Single-Root Input/Output Virtualization (SR-IOV) capabilities.
// }

// The capabilities of a controller.
type ControllerCapabilities struct {
	NetworkPortCount int `json:"NetworkPortCount"` // The number of physical ports on this controller.
	// NetworkDeviceFunctionCount int                   // The maximum number of physical functions available on this controller.
	// DataCenterBridging         DataCenterBridging    // Data Center Bridging (DCB) for this controller.
	// VirtualizationOffload      VirtualizationOffload // Virtualization offload for this controller.
	// NPIV                       NPIV                  // N_Port ID Virtualization (NPIV) capabilties for this controller.
}

// A network controller ASIC that makes up part of a NetworkAdapter.
type Controller struct {
	FirmwarePackageVersion string                 `json:"FirmwarePackageVersion"` // The version of the user-facing firmware package.
	ControllerCapabilities ControllerCapabilities `json:"Slot"`                   // The capabilities of this controller.
	NetworkPorts           []NetworkPort          `json:"NetworkPorts"`           // Contains the members of this collection.
}

func (dto *Controller) Load(m *model.Controller) {
	dto.FirmwarePackageVersion = m.FirmwarePackageVersion
	dto.ControllerCapabilities.NetworkPortCount = m.ControllerCapabilities.NetworkPortCount
	for i, _ := range m.NetworkPorts {
		portD := NetworkPort{}
		portD.Load(&m.NetworkPorts[i])
		dto.NetworkPorts = append(dto.NetworkPorts, portD)
	}
}

// A NetworkAdapter represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	ResourceResponse
	ProductInfoResponse
	Controllers []Controller `json:"Controllers"` // The set of network controllers ASICs that make up this NetworkAdapter.
	//	NetworkPorts []NetworkPort // Contains the members of this collection.
	//	NetworkDeviceFunctions []NetworkDeviceFunction // Contains the members of this collection.
}

func (dto *NetworkAdapter) Load(m *model.NetworkAdapter) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	for i, _ := range m.Controllers {
		controllerD := Controller{}
		controllerD.Load(&m.Controllers[i])
		dto.Controllers = append(dto.Controllers, controllerD)
	}
}
