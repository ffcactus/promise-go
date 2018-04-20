package dto

type PCIeFunctionLinks struct {
	EthernetInterfaces []ResourceRef // An array of references to the ethernet interfaces which the PCIe device produces.
	Drives             []ResourceRef // An array of references to the drives which the PCIe device produces.
	StorageControllers []ResourceRef // An array of references to the storage controllers which the PCIe device produces.
}

// This is the schema definition for the PCIeFunction resource.  It represents the properties of a PCIeFunction attached to a System.
type GetPCIeFunctionResponse struct {
	Resource
	Links             PCIeFunctionLinks
	DeviceClass       *string // The class for this PCIe Function.
	DeviceID          *string // The Device ID of this PCIe function.
	VendorID          *string // The Vendor ID of this PCIe function.
	SubsystemID       *string // The Subsystem ID of this PCIe function.
	SubsystemVendorID *string // The Subsystem Vendor ID of this PCIe function.
}
