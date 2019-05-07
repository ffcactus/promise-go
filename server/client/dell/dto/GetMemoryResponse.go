package dto

// MemoryLocation is memory connection information to sockets and memory controllers.
type MemoryLocation struct {
	Socket     *int // Socket number in which Memory is connected.
	Controller *int // Memory controller number in which Memory is connected.
	Channel    *int // Channel number in which Memory is connected.
	Slot       *int // Slot number in which Memory is connected.
}

// GetMemoryResponse definition for definition of a Memory and its configuration.
type GetMemoryResponse struct {
	Resource
	ProductInfo
	MemoryType        *string         `json:"MemoryType"`        // The type of Memory.
	MemoryDeviceType  *string         `json:"MemoryDeviceType"`  // Type details of the Memory.
	CapacityMiB       *int            `json:"CapacityMiB"`       // Memory Capacity in MiB.
	DataWidthBits     *int            `json:"DataWidthBits"`     // Data Width in bits.
	BusWidthBits      *int            `json:"BusWidthBits"`      // Bus Width in bits.
	VendorID          *string         `json:"VendorID"`          // Vendor ID.
	DeviceID          *string         `json:"DeviceID"`          // Device ID.
	SubsystemVendorID *string         `json:"SubsystemVendorID"` // SubSystem Vendor ID.
	SubsystemDeviceID *string         `json:"SubsystemDeviceID"` // Subsystem Device ID.
	RankCount         *int            `json:"RankCount"`         // Number of ranks available in the Memory.
	DeviceLocator     *string         `json:"DeviceLocator"`     // Location of the Memory in the platform.
	MemoryLocation    *MemoryLocation `json:"MemoryLocation"`    // Memory connection information to sockets and memory controllers.
	ErrorCorrection   *string         `json:"ErrorCorrection"`   // Error correction scheme supported for this memory.
	OperatingSpeedMhz *int            `json:"OperatingSpeedMhz"` // Operating speed of Memory in MHz.
}
