package entity

import (
	"promise/server/object/model"
)

// PCIeDevice This is the schema definition for the PCIeDevice resource.  It represents the properties of a PCIeDevice attached to a System.
type PCIeDevice struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	DeviceType      *string        // The device type for this PCIe device.
	FirmwareVersion *string        // The version of firmware for this PCIe device.
	PCIeFunctions   []PCIeFunction `gorm:"ForeignKey:PCIeDeviceRef"`
}

// ToModel will create a new model from entity.
func (e *PCIeDevice) ToModel() *model.PCIeDevice {
	m := new(model.PCIeDevice)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.DeviceType = e.DeviceType
	m.FirmwareVersion = e.FirmwareVersion
	for i := range e.PCIeFunctions {
		m.PCIeFunctions = append(m.PCIeFunctions, *e.PCIeFunctions[i].ToModel())
	}
	return m
}
