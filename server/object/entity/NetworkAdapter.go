package entity

import (
	commonUtil "promise/common/util"
	"promise/server/object/model"
)

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

// ToModel will create a new model from entity.
func (e *NetworkAdapter) ToModel() *model.NetworkAdapter {
	m := new(model.NetworkAdapter)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	for i := range e.Controllers {
		controllerE := e.Controllers[i]
		controllerM := model.Controller{}
		controllerM.FirmwarePackageVersion = controllerE.FirmwarePackageVersion
		controllerM.ControllerCapabilities.NetworkPortCount = controllerE.ControllerCapabilitiesNetworkPortCount
		for j := range controllerE.NetworkPorts {
			portE := controllerE.NetworkPorts[j]
			portM := model.NetworkPort{}
			portM.PhysicalPortNumber = portE.PhysicalPortNumber
			portM.LinkStatus = portE.LinkStatus
			a := []string{}
			commonUtil.StringToStruct(portE.AssociatedNetworkAddresses, &a)
			portM.AssociatedNetworkAddresses = a
			controllerM.NetworkPorts = append(controllerM.NetworkPorts, portM)
		}
		m.Controllers = append(m.Controllers, controllerM)
	}
	return m
}
