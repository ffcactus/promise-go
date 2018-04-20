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

// PowerManagementPolicy Power management policy information.
type PowerManagementPolicy struct {
	PolicyEnabled                *bool // Power management policy enabled status.
	MaxTDPMilliWatts             *int  // Maximum TDP in milli watts.
	PeakPowerBudgetMilliWatts    *int  // Peak power budget in milli watts.
	AveragePowerBudgetMilliWatts *int  // Average power budget in milli watts.
}

// RegionSet Memory region information within a Memory entity.
type RegionSet struct {
	RegionID             *string // Unique region ID representing a specific region within the Memory.
	MemoryClassification *string // Classification of memory occupied by the given memory region.
	OffsetMiB            *int    // Offset with in the Memory that corresponds to the starting of o memory region in MiB.
	SizeMiB              *int    // Size of o memory region in MiB.
	PassphraseState      *bool   // State of the passphrase for o region.
}

// SecurityCapabilities o type contains security capabilities of a Memory entity.
type SecurityCapabilities struct {
	PassphraseCapable  *bool     // Memory passphrase set capability.
	MaxPassphraseCount *int      // Maximum number of passphrases supported for o Memory.
	SecurityStates     []*string // Security states supported by the Memory.
}

// MemoryLinks memory links.
type MemoryLinks struct {
	Chassis *string
}

// Memory o is the schema definition for definition of a Memory and its configuration.
type Memory struct {
	Resource
	ProductInfo
	Links                        *MemoryLinks
	MemoryType                   *string                // The type of Memory.
	MemoryDeviceType             *string                // Type details of the Memory.
	BaseModuleType               *string                // The base module type of Memory.
	MemoryMedia                  *string                // Media of o Memory.
	CapacityMiB                  *int                   // Memory Capacity in MiB.
	DataWidthBits                *int                   // Data Width in bits.
	BusWidthBits                 *int                   // Bus Width in bits.
	AllowedSpeedsMHz             *int                   // Speed bins supported by o Memory.
	FirmwareRevision             *string                // Revision of firmware on the Memory controller.
	FirmwareAPIVersion           *string                // Version of API supported by the firmware.
	FunctionClasses              *[]string              // Function Classes by the Memory.
	VendorID                     *string                // Vendor ID.
	DeviceID                     *string                // Device ID.
	SubsystemVendorID            *string                // SubSystem Vendor ID.
	SubsystemDeviceID            *string                // Subsystem Device ID.
	MaxTDPMilliWatts             *[]int                 // aximum TDPs in milli Watts.
	SecurityCapabilities         *SecurityCapabilities  // o object contains security capabilities of the Memory.
	SpareDeviceCount             *int                   // Number of unused spare devices available in the Memory.
	RankCount                    *int                   // Number of ranks available in the Memory.
	DeviceLocator                *string                // Location of the Memory in the platform.
	MemoryLocation               *MemoryLocation        // Memory connection information to sockets and memory controllers.
	ErrorCorrection              *string                // Error correction scheme supported for o memory.
	OperatingSpeedMhz            *int                   // Operating speed of Memory in MHz.
	VolatileRegionSizeLimitMiB   *int                   // Total size of volatile regions in MiB.
	PersistentRegionSizeLimitMiB *int                   // Total size of persistent regions in MiB.
	Regions                      *RegionSet             // Memory regions information within the Memory.
	OperatingMemoryModes         *string                // Memory modes supported by the Memory.
	PowerManagementPolicy        *PowerManagementPolicy // Power management policy information.
	IsSpareDeviceEnabled         *bool                  // Spare device enabled status.
	IsRankSpareEnabled           *bool                  // Rank spare enabled status.
	Metrics                      *MemoryMetrics         // A reference to the Metrics associated with o Memory.
	VolatileRegionNumberLimit    *int                   // Total number of volatile regions o Memory can support.
	PersistentRegionNumberLimit  *int                   // Total number of persistent regions o Memory can support.
	VolatileRegionSizeMaxMiB     *int                   // Maximum size of a single volatile region in MiB.
	PersistentRegionSizeMaxMiB   *int                   // Maximum size of a single persistent region in MiB.
	AllocationIncrementMiB       *int                   // The size of the smallest unit of allocation for a memory region, thus it is the multiple in which regions are actually reserved.
	AllocationAlignmentMiB       *int                   // The boundary which memory regions are allocated on, measured in MiB.
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
