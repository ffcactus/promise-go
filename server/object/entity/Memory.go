package entity

import (
	"promise/server/object/model"
)

// Memory This is the schema definition for definition of a Memory and its configuration.
type Memory struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	LinksChassis      *string
	MemoryType        *string // The type of Memory.
	MemoryDeviceType  *string // Type details of the Memory.
	CapacityMiB       *int    // Memory Capacity in MiB.
	DataWidthBits     *int    // Data Width in bits.
	BusWidthBits      *int    // Bus Width in bits.
	VendorID          *string // Vendor ID.
	DeviceID          *string // Device ID.
	SubsystemVendorID *string // SubSystem Vendor ID.
	SubsystemDeviceID *string // Subsystem Device ID.
	RankCount         *int    // Number of ranks available in the Memory.
	DeviceLocator     *string // Location of the Memory in the platform.
	ErrorCorrection   *string // Error correction scheme supported for this memory.
	OperatingSpeedMhz *int    // Operating speed of Memory in MHz.

	// expand MemoryLocation
	MemoryLocationSocket     *int // Socket number in which Memory is connected.
	MemoryLocationController *int // Memory controller number in which Memory is connected.
	MemoryLocationChannel    *int // Channel number in which Memory is connected.
	MemoryLocationSlot       *int // Slot number in which Memory is connected.
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

// Load will load data from model.
func (e *Memory) Load(m *model.Memory) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	updateProductInfoEntity(&(*e).ProductInfo, &(*m).ProductInfo)
	e.MemoryType = m.MemoryType
	e.MemoryDeviceType = m.MemoryDeviceType
	e.CapacityMiB = m.CapacityMiB
	e.DataWidthBits = m.DataWidthBits
	e.BusWidthBits = m.BusWidthBits
	e.VendorID = m.VendorID
	e.DeviceID = m.DeviceID
	e.SubsystemVendorID = m.SubsystemVendorID
	e.SubsystemDeviceID = m.SubsystemDeviceID
	e.RankCount = m.RankCount
	e.DeviceLocator = m.DeviceLocator
	if m.MemoryLocation != nil {
		e.MemoryLocationSocket = m.MemoryLocation.Socket
		e.MemoryLocationController = m.MemoryLocation.Controller
		e.MemoryLocationChannel = m.MemoryLocation.Channel
		e.MemoryLocationSlot = m.MemoryLocation.Slot
	}
	e.ErrorCorrection = m.ErrorCorrection
	e.OperatingSpeedMhz = m.OperatingSpeedMhz
}
