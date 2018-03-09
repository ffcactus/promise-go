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

func (dto *PCIeDevice) Load(m *model.PCIeDevice, ethernetInterfaces []model.EthernetInterface) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.DeviceType = m.DeviceType
	dto.FirmwareVersion = m.FirmwareVersion
	for i := range m.PCIeFunctions {
		each := new(PCIeFunction)
		each.Load(&m.PCIeFunctions[i], ethernetInterfaces)
		dto.PCIeFunctions = append(dto.PCIeFunctions, *each)
	}
}
