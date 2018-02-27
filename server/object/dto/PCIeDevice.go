package dto

import (
	"promise/server/object/model"
)

// This is the schema definition for the PCIeDevice resource.  It represents the properties of a PCIeDevice attached to a System.
type PCIeDevice struct {
	ResourceResponse
	ProductInfoResponse
	DeviceType      *string // The device type for this PCIe device.
	FirmwareVersion *string // The version of firmware for this PCIe device.
	PCIeFunctions   []PCIeFunction
}

func (this *PCIeDevice) Load(m *model.PCIeDevice, ethernetInterfaces []model.EthernetInterface) {
	this.LoadResourceResponse(&m.Resource)
	this.LoadProductInfoResponse(&m.ProductInfo)
	this.DeviceType = m.DeviceType
	this.FirmwareVersion = m.FirmwareVersion
	for i, _ := range m.PCIeFunctions {
		each := new(PCIeFunction)
		each.Load(&m.PCIeFunctions[i], ethernetInterfaces)
		this.PCIeFunctions = append(this.PCIeFunctions, *each)
	}
}
