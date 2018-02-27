package model

// PCIeDevice This is the schema definition for the PCIeDevice resource.  It represents the properties of a PCIeDevice attached to a System.
type PCIeDevice struct {
	Resource
	ProductInfo
	DeviceType      *string // The device type for this PCIe device.
	FirmwareVersion *string // The version of firmware for this PCIe device.
	PCIeFunctions   []PCIeFunction
}
