package entity

// PCIeFunction This is the schema definition for the PCIeFunction resource.  It represents the properties of a PCIeFunction attached to a System.
type PCIeFunction struct {
	PCIeDeviceRef uint
	EmbeddedResource
	DeviceClass        *string // The class for this PCIe Function.
	DeviceID           *string // The Device ID of this PCIe function.
	VendorID           *string // The Vendor ID of this PCIe function.
	SubsystemID        *string // The Subsystem ID of this PCIe function.
	SubsystemVendorID  *string // The Subsystem Vendor ID of this PCIe function.
	EthernetInterfaces string  // An array of references to the ethernet interfaces which the PCIe device produces.
}
