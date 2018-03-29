package entity

import (
	"promise/server/object/model"
)

// Memory This is the schema definition for definition of a Memory and its configuration.
type Memory struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	LinksChassis                 *string
	MemoryType                   *string // The type of Memory.
	MemoryDeviceType             *string // Type details of the Memory.
	BaseModuleType               *string // The base module type of Memory.
	MemoryMedia                  *string // Media of this Memory.
	CapacityMiB                  *int    // Memory Capacity in MiB.
	DataWidthBits                *int    // Data Width in bits.
	BusWidthBits                 *int    // Bus Width in bits.
	AllowedSpeedsMHz             *int    // Speed bins supported by this Memory.
	FirmwareRevision             *string // Revision of firmware on the Memory controller.
	FirmwareAPIVersion           *string // Version of API supported by the firmware.
	VendorID                     *string // Vendor ID.
	DeviceID                     *string // Device ID.
	SubsystemVendorID            *string // SubSystem Vendor ID.
	SubsystemDeviceID            *string // Subsystem Device ID.
	SpareDeviceCount             *int    // Number of unused spare devices available in the Memory.
	RankCount                    *int    // Number of ranks available in the Memory.
	DeviceLocator                *string // Location of the Memory in the platform.
	MemoryLocationSocket         *int    // Memory connection information to sockets and memory controllers.
	MemoryLocationController     *int    // Memory connection information to sockets and memory controllers.
	MemoryLocationChannel        *int    // Memory connection information to sockets and memory controllers.
	MemoryLocationSlot           *int    // Memory connection information to sockets and memory controllers.
	ErrorCorrection              *string // Error correction scheme supported for this memory.
	OperatingSpeedMhz            *int    // Operating speed of Memory in MHz.
	VolatileRegionSizeLimitMiB   *int    // Total size of volatile regions in MiB.
	PersistentRegionSizeLimitMiB *int    // Total size of persistent regions in MiB.
	OperatingMemoryModes         *string // Memory modes supported by the Memory.
	IsSpareDeviceEnabled         *bool   // Spare device enabled status.
	IsRankSpareEnabled           *bool   // Rank spare enabled status.
	VolatileRegionNumberLimit    *int    // Total number of volatile regions this Memory can support.
	PersistentRegionNumberLimit  *int    // Total number of persistent regions this Memory can support.
	VolatileRegionSizeMaxMiB     *int    // Maximum size of a single volatile region in MiB.
	PersistentRegionSizeMaxMiB   *int    // Maximum size of a single persistent region in MiB.
	AllocationIncrementMiB       *int    // The size of the smallest unit of allocation for a memory region, thus it is the multiple in which regions are actually reserved.
	AllocationAlignmentMiB       *int    // The boundary which memory regions are allocated on, measured in MiB.
}

// ToModel will create a new model from entity.
func (e *Memory) ToModel() *model.Memory {
	m := model.Memory{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.CapacityMiB = e.CapacityMiB
	m.OperatingSpeedMhz = e.OperatingSpeedMhz
	m.MemoryDeviceType = e.MemoryDeviceType
	m.DataWidthBits = e.DataWidthBits
	m.RankCount = e.RankCount
	m.DeviceLocator = e.DeviceLocator
	if e.MemoryLocationSocket != nil ||
		e.MemoryLocationController != nil ||
		e.MemoryLocationChannel != nil ||
		e.MemoryLocationSlot != nil {
		m.MemoryLocation = new(model.MemoryLocation)
		m.MemoryLocation.Socket = e.MemoryLocationSocket
		m.MemoryLocation.Controller = e.MemoryLocationController
		m.MemoryLocation.Channel = e.MemoryLocationChannel
		m.MemoryLocation.Slot = e.MemoryLocationSlot
	}
	return &m
}
