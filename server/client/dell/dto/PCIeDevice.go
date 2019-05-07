package dto

type PCIeDeviceLinks struct {
	PCIeFunctions []ResourceRef // An array of references to the ethernet interfaces which the PCIe device produces.
}

// This is the schema definition for the PCIeDevice resource.  It represents the properties of a PCIeDevice attached to a System.
type GetPCIeDeviceResponse struct {
	Resource
	ProductInfo
	DeviceType      *string // The device type for this PCIe device.
	FirmwareVersion *string // The version of firmware for this PCIe device.
	Links           PCIeDeviceLinks
}
