package model

import (
	log "github.com/sirupsen/logrus"
)

// MemoryLocation Memory connection information to sockets and memory controllers.
type MemoryLocation struct {
	Socket     *int // Socket number in which Memory is connected.
	Controller *int // Memory controller number in which Memory is connected.
	Channel    *int // Channel number in which Memory is connected.
	Slot       *int // Slot number in which Memory is connected.
}

// Memory o is the schema definition for definition of a Memory and its configuration.
type Memory struct {
	Resource
	ProductInfo
	MemoryType        *string         // The type of Memory.
	MemoryDeviceType  *string         // Type details of the Memory.
	CapacityMiB       *int            // Memory Capacity in MiB.
	DataWidthBits     *int            // Data Width in bits.
	BusWidthBits      *int            // Bus Width in bits.
	VendorID          *string         // Vendor ID.
	DeviceID          *string         // Device ID.
	SubsystemVendorID *string         // SubSystem Vendor ID.
	SubsystemDeviceID *string         // Subsystem Device ID.
	RankCount         *int            // Number of ranks available in the Memory.
	DeviceLocator     *string         // Location of the Memory in the platform.
	MemoryLocation    *MemoryLocation // Memory connection information to sockets and memory controllers.
	ErrorCorrection   *string         // Error correction scheme supported for o memory.
	OperatingSpeedMhz *int            // Operating speed of Memory in MHz.
}

// Print print the memory.
func (o *Memory) Print() {
	if o.URI != nil {
		log.Info("URI = %#v\n", *o.URI)
	} else {
		log.Info("URI = nil\n")
	}
	if o.PhysicalState != nil {
		log.Info("PhysicalState = %#v\n", *o.PhysicalState)
	} else {
		log.Info("PhysicalState = nil\n")
	}
	if o.Manufacturer != nil {
		log.Info("Manufacturer = %#v\n", *o.Manufacturer)
	} else {
		log.Info("Manufacturer = nil\n")
	}
	if o.MemoryLocation != nil {
		if o.MemoryLocation.Socket != nil {
			log.Info("MemoryLocation.Socket = %#v\n", *o.MemoryLocation.Socket)
		} else {
			log.Info("MemoryLocation.Socket = nil\n")
		}
	} else {
		log.Info("MemoryLocation = nil\n")
	}

}
