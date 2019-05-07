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
	MemoryType        *string         `json:"MemoryType,omitempty"`        // The type of Memory.
	MemoryDeviceType  *string         `json:"MemoryDeviceType,omitempty"`  // Type details of the Memory.
	CapacityMiB       *int            `json:"CapacityMiB,omitempty"`       // Memory Capacity in MiB.
	DataWidthBits     *int            `json:"DataWidthBits,omitempty"`     // Data Width in bits.
	BusWidthBits      *int            `json:"BusWidthBits,omitempty"`      // Bus Width in bits.
	VendorID          *string         `json:"VendorID,omitempty"`          // Vendor Idto.
	DeviceID          *string         `json:"DeviceID,omitempty"`          // Device Idto.
	SubsystemVendorID *string         `json:"SubsystemVendorID,omitempty"` // SubSystem Vendor Idto.
	SubsystemDeviceID *string         `json:"SubsystemDeviceID,omitempty"` // Subsystem Device Idto.
	RankCount         *int            `json:"RankCount,omitempty"`         // Number of ranks available in the Memory.
	DeviceLocator     *string         `json:"DeviceLocator,omitempty"`     // Location of the Memory in the platform.
	MemoryLocation    *MemoryLocation `json:"MemoryLocation,omitempty"`    // Memory connection information to sockets and memory controllers.
	ErrorCorrection   *string         `json:"ErrorCorrection,omitempty"`   // Error correction scheme supported for dto memory.
	OperatingSpeedMhz *int            `json:"OperatingSpeedMhz,omitempty"` // Operating speed of Memory in MHz.
}

// Load will load data from model.
func (dto *Memory) Load(m *model.Memory) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.MemoryType = m.MemoryType
	dto.MemoryDeviceType = m.MemoryDeviceType
	dto.CapacityMiB = m.CapacityMiB
	dto.DataWidthBits = m.DataWidthBits
	dto.BusWidthBits = m.BusWidthBits
	dto.VendorID = m.VendorID
	dto.DeviceID = m.DeviceID
	dto.SubsystemVendorID = m.SubsystemVendorID
	dto.SubsystemDeviceID = m.SubsystemDeviceID
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
}
