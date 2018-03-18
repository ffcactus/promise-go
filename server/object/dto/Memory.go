package dto

import (
	"promise/server/object/model"
)

// MemoryLocation is DTO.
type MemoryLocation struct {
	Socket     *int `json:"Socket,omitempty"`
	Controller *int `json:"Controller,omitempty"`
	Channel    *int `json:"Channel,omitempty"`
	Slot       *int `json:"Slot,omitempty"`
}

// Memory is DTO.
type Memory struct {
	ResourceResponse
	ProductInfoResponse
	MemoryType                   *string         `json:"MemoryType,omitempty"`                   // The type of Memory.
	MemoryDeviceType             *string         `json:"MemoryDeviceType,omitempty"`             // Type details of the Memory.
	BaseModuleType               *string         `json:"BaseModuleType,omitempty"`               // The base module type of Memory.
	MemoryMedia                  *string         `json:"MemoryMedia,omitempty"`                  // Media of dto Memory.
	CapacityMiB                  *int            `json:"CapacityMiB,omitempty"`                  // Memory Capacity in MiB.
	DataWidthBits                *int            `json:"DataWidthBits,omitempty"`                // Data Width in bits.
	BusWidthBits                 *int            `json:"BusWidthBits,omitempty"`                 // Bus Width in bits.
	AllowedSpeedsMHz             *int            `json:"AllowedSpeedsMHz,omitempty"`             // Speed bins supported by dto Memory.
	FirmwareRevision             *string         `json:"FirmwareRevision,omitempty"`             // Revision of firmware on the Memory controller.
	FirmwareAPIVersion           *string         `json:"FirmwareAPIVersion,omitempty"`           // Version of API supported by the firmware.
	VendorID                     *string         `json:"VendorID,omitempty"`                     // Vendor Idto.
	DeviceID                     *string         `json:"DeviceID,omitempty"`                     // Device Idto.
	SubsystemVendorID            *string         `json:"SubsystemVendorID,omitempty"`            // SubSystem Vendor Idto.
	SubsystemDeviceID            *string         `json:"SubsystemDeviceID,omitempty"`            // Subsystem Device Idto.
	SpareDeviceCount             *int            `json:"SpareDeviceCount,omitempty"`             // Number of unused spare devices available in the Memory.
	RankCount                    *int            `json:"RankCount,omitempty"`                    // Number of ranks available in the Memory.
	DeviceLocator                *string         `json:"DeviceLocator,omitempty"`                // Location of the Memory in the platform.
	MemoryLocation               *MemoryLocation `json:"MemoryLocation,omitempty"`               // Memory connection information to sockets and memory controllers.
	ErrorCorrection              *string         `json:"ErrorCorrection,omitempty"`              // Error correction scheme supported for dto memory.
	OperatingSpeedMhz            *int            `json:"OperatingSpeedMhz,omitempty"`            // Operating speed of Memory in MHz.
	VolatileRegionSizeLimitMiB   *int            `json:"VolatileRegionSizeLimitMiB,omitempty"`   // Total size of volatile regions in MiB.
	PersistentRegionSizeLimitMiB *int            `json:"PersistentRegionSizeLimitMiB,omitempty"` // Total size of persistent regions in MiB.
	OperatingMemoryModes         *string         `json:"OperatingMemoryModes,omitempty"`         // Memory modes supported by the Memory.
	IsSpareDeviceEnabled         *bool           `json:"IsSpareDeviceEnabled,omitempty"`         // Spare device enabled status.
	IsRankSpareEnabled           *bool           `json:"IsRankSpareEnabled,omitempty"`           // Rank spare enabled status.
	VolatileRegionNumberLimit    *int            `json:"VolatileRegionNumberLimit,omitempty"`    // Total number of volatile regions dto Memory can support.
	PersistentRegionNumberLimit  *int            `json:"PersistentRegionNumberLimit,omitempty"`  // Total number of persistent regions dto Memory can support.
	VolatileRegionSizeMaxMiB     *int            `json:"VolatileRegionSizeMaxMiB,omitempty"`     // Maximum size of a single volatile region in MiB.
	PersistentRegionSizeMaxMiB   *int            `json:"PersistentRegionSizeMaxMiB,omitempty"`   // Maximum size of a single persistent region in MiB.
	AllocationIncrementMiB       *int            `json:"AllocationIncrementMiB,omitempty"`       // The size of the smallest unit of allocation for a memory region, thus it is the multiple in which regions are actually reservedto.
	AllocationAlignmentMiB       *int            `json:"AllocationAlignmentMiB,omitempty"`       // The boundary which memory regions are allocated on, measured in MiB.
}

// Load will load data from model.
func (dto *Memory) Load(m *model.Memory) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.MemoryType = m.MemoryType
	dto.MemoryDeviceType = m.MemoryDeviceType
	dto.BaseModuleType = m.BaseModuleType
	dto.MemoryMedia = m.MemoryMedia
	dto.CapacityMiB = m.CapacityMiB
	dto.DataWidthBits = m.DataWidthBits
	dto.BusWidthBits = m.BusWidthBits
	dto.AllowedSpeedsMHz = m.AllowedSpeedsMHz

	dto.FirmwareRevision = m.FirmwareRevision
	dto.FirmwareAPIVersion = m.FirmwareAPIVersion
	dto.VendorID = m.VendorID
	dto.DeviceID = m.DeviceID
	dto.SubsystemVendorID = m.SubsystemVendorID
	dto.SubsystemDeviceID = m.SubsystemDeviceID
	dto.SpareDeviceCount = m.SpareDeviceCount
	dto.DeviceLocator = m.DeviceLocator
	if m.MemoryLocation != nil {
		dto.MemoryLocation = new(MemoryLocation)
		dto.MemoryLocation.Socket = m.MemoryLocation.Socket
		dto.MemoryLocation.Controller = m.MemoryLocation.Controller
		dto.MemoryLocation.Channel = m.MemoryLocation.Channel
		dto.MemoryLocation.Slot = m.MemoryLocation.Slot
	}
	dto.ErrorCorrection = m.ErrorCorrection
	dto.OperatingSpeedMhz = m.OperatingSpeedMhz
	dto.VolatileRegionSizeLimitMiB = m.VolatileRegionSizeLimitMiB
	dto.PersistentRegionSizeLimitMiB = m.PersistentRegionSizeLimitMiB
	dto.OperatingMemoryModes = m.OperatingMemoryModes
	dto.IsSpareDeviceEnabled = m.IsSpareDeviceEnabled
	dto.IsRankSpareEnabled = m.IsRankSpareEnabled
	dto.VolatileRegionNumberLimit = m.VolatileRegionNumberLimit
	dto.VolatileRegionSizeMaxMiB = m.VolatileRegionSizeMaxMiB
	dto.AllocationIncrementMiB = m.AllocationIncrementMiB
}
