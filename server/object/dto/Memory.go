package dto

import (
	"promise/server/object/model"
)

type MemoryLocation struct {
	Socket     *int `json:"Socket,omitempty"`
	Controller *int `json:"Controller,omitempty"`
	Channel    *int `json:"Channel,omitempty"`
	Slot       *int `json:"Slot,omitempty"`
}

type Memory struct {
	ResourceResponse
	ProductInfoResponse
	MemoryType                   *string         `json:"MemoryType,omitempty"`                   // The type of Memory.
	MemoryDeviceType             *string         `json:"MemoryDeviceType,omitempty"`             // Type details of the Memory.
	BaseModuleType               *string         `json:"BaseModuleType,omitempty"`               // The base module type of Memory.
	MemoryMedia                  *string         `json:"MemoryMedia,omitempty"`                  // Media of this Memory.
	CapacityMiB                  *int            `json:"CapacityMiB,omitempty"`                  // Memory Capacity in MiB.
	DataWidthBits                *int            `json:"DataWidthBits,omitempty"`                // Data Width in bits.
	BusWidthBits                 *int            `json:"BusWidthBits,omitempty"`                 // Bus Width in bits.
	AllowedSpeedsMHz             *int            `json:"AllowedSpeedsMHz,omitempty"`             // Speed bins supported by this Memory.
	FirmwareRevision             *string         `json:"FirmwareRevision,omitempty"`             // Revision of firmware on the Memory controller.
	FirmwareAPIVersion           *string         `json:"FirmwareAPIVersion,omitempty"`           // Version of API supported by the firmware.
	VendorID                     *string         `json:"VendorID,omitempty"`                     // Vendor Ithis.
	DeviceID                     *string         `json:"DeviceID,omitempty"`                     // Device Ithis.
	SubsystemVendorID            *string         `json:"SubsystemVendorID,omitempty"`            // SubSystem Vendor Ithis.
	SubsystemDeviceID            *string         `json:"SubsystemDeviceID,omitempty"`            // Subsystem Device Ithis.
	SpareDeviceCount             *int            `json:"SpareDeviceCount,omitempty"`             // Number of unused spare devices available in the Memory.
	RankCount                    *int            `json:"RankCount,omitempty"`                    // Number of ranks available in the Memory.
	DeviceLocator                *string         `json:"DeviceLocator,omitempty"`                // Location of the Memory in the platform.
	MemoryLocation               *MemoryLocation `json:"MemoryLocation,omitempty"`               // Memory connection information to sockets and memory controllers.
	ErrorCorrection              *string         `json:"ErrorCorrection,omitempty"`              // Error correction scheme supported for this memory.
	OperatingSpeedMhz            *int            `json:"OperatingSpeedMhz,omitempty"`            // Operating speed of Memory in MHz.
	VolatileRegionSizeLimitMiB   *int            `json:"VolatileRegionSizeLimitMiB,omitempty"`   // Total size of volatile regions in MiB.
	PersistentRegionSizeLimitMiB *int            `json:"PersistentRegionSizeLimitMiB,omitempty"` // Total size of persistent regions in MiB.
	OperatingMemoryModes         *string         `json:"OperatingMemoryModes,omitempty"`         // Memory modes supported by the Memory.
	IsSpareDeviceEnabled         *bool           `json:"IsSpareDeviceEnabled,omitempty"`         // Spare device enabled status.
	IsRankSpareEnabled           *bool           `json:"IsRankSpareEnabled,omitempty"`           // Rank spare enabled status.
	VolatileRegionNumberLimit    *int            `json:"VolatileRegionNumberLimit,omitempty"`    // Total number of volatile regions this Memory can support.
	PersistentRegionNumberLimit  *int            `json:"PersistentRegionNumberLimit,omitempty"`  // Total number of persistent regions this Memory can support.
	VolatileRegionSizeMaxMiB     *int            `json:"VolatileRegionSizeMaxMiB,omitempty"`     // Maximum size of a single volatile region in MiB.
	PersistentRegionSizeMaxMiB   *int            `json:"PersistentRegionSizeMaxMiB,omitempty"`   // Maximum size of a single persistent region in MiB.
	AllocationIncrementMiB       *int            `json:"AllocationIncrementMiB,omitempty"`       // The size of the smallest unit of allocation for a memory region, thus it is the multiple in which regions are actually reservethis.
	AllocationAlignmentMiB       *int            `json:"AllocationAlignmentMiB,omitempty"`       // The boundary which memory regions are allocated on, measured in MiB.
}

func (this *Memory) Load(m *model.Memory) {
	this.LoadResourceResponse(&m.Resource)
	this.LoadProductInfoResponse(&m.ProductInfo)
	this.MemoryType = m.MemoryType
	this.MemoryDeviceType = m.MemoryDeviceType
	this.BaseModuleType = m.BaseModuleType
	this.MemoryMedia = m.MemoryMedia
	this.CapacityMiB = m.CapacityMiB
	this.DataWidthBits = m.DataWidthBits
	this.BusWidthBits = m.BusWidthBits
	this.AllowedSpeedsMHz = m.AllowedSpeedsMHz

	this.FirmwareRevision = m.FirmwareRevision
	this.FirmwareAPIVersion = m.FirmwareAPIVersion
	this.VendorID = m.VendorID
	this.DeviceID = m.DeviceID
	this.SubsystemVendorID = m.SubsystemVendorID
	this.SubsystemDeviceID = m.SubsystemDeviceID
	this.SpareDeviceCount = m.SpareDeviceCount
	this.DeviceLocator = m.DeviceLocator
	if m.MemoryLocation != nil {
		this.MemoryLocation = new(MemoryLocation)
		this.MemoryLocation.Socket = m.MemoryLocation.Socket
		this.MemoryLocation.Controller = m.MemoryLocation.Controller
		this.MemoryLocation.Channel = m.MemoryLocation.Channel
		this.MemoryLocation.Slot = m.MemoryLocation.Slot
	}
	this.ErrorCorrection = m.ErrorCorrection
	this.OperatingSpeedMhz = m.OperatingSpeedMhz
	this.VolatileRegionSizeLimitMiB = m.VolatileRegionSizeLimitMiB
	this.PersistentRegionSizeLimitMiB = m.PersistentRegionSizeLimitMiB
	this.OperatingMemoryModes = m.OperatingMemoryModes
	this.IsSpareDeviceEnabled = m.IsSpareDeviceEnabled
	this.IsRankSpareEnabled = m.IsRankSpareEnabled
	this.VolatileRegionNumberLimit = m.VolatileRegionNumberLimit
	this.VolatileRegionSizeMaxMiB = m.VolatileRegionSizeMaxMiB
	this.AllocationIncrementMiB = m.AllocationIncrementMiB
}
