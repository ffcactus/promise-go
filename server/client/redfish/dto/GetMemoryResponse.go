package dto

// Memory connection information to sockets and memory controllers.
type MemoryLocation struct {
	Socket     *int // Socket number in which Memory is connected.
	Controller *int // Memory controller number in which Memory is connected.
	Channel    *int // Channel number in which Memory is connected.
	Slot       *int // Slot number in which Memory is connected.
}

// Power management policy information.
type PowerManagementPolicy struct {
	PolicyEnabled                *bool // Power management policy enabled status.
	MaxTDPMilliWatts             *int  // Maximum TDP in milli watts.
	PeakPowerBudgetMilliWatts    *int  // Peak power budget in milli watts.
	AveragePowerBudgetMilliWatts *int  // Average power budget in milli watts.
}

// Memory region information within a Memory entity.
type RegionSet struct {
	RegionId             *string // Unique region ID representing a specific region within the Memory.
	MemoryClassification *string // Classification of memory occupied by the given memory region.
	OffsetMiB            *int    // Offset with in the Memory that corresponds to the starting of this memory region in MiB.
	SizeMiB              *int    // Size of this memory region in MiB.
	PassphraseState      *bool   // State of the passphrase for this region.
}

// This type contains security capabilities of a Memory entity.
type SecurityCapabilities struct {
	PassphraseCapable  *bool     // Memory passphrase set capability.
	MaxPassphraseCount *int      // Maximum number of passphrases supported for this Memory.
	SecurityStates     []*string // Security states supported by the Memory.
}

type MemoryLinks struct {
	Chassis *string
}

// This is the schema definition for definition of a Memory and its configuration.
type GetMemoryResponse struct {
	Resource
	ProductInfo
	Links                        *MemoryLinks           `json:"Links"`
	MemoryType                   *string                `json:"MemoryType"`                   // The type of Memory.
	MemoryDeviceType             *string                `json:"MemoryDeviceType"`             // Type details of the Memory.
	BaseModuleType               *string                `json:"BaseModuleType"`               // The base module type of Memory.
	MemoryMedia                  *string                `json:"MemoryMedia"`                  // Media of this Memory.
	CapacityMiB                  *int                   `json:"CapacityMiB"`                  // Memory Capacity in MiB.
	DataWidthBits                *int                   `json:"DataWidthBits"`                // Data Width in bits.
	BusWidthBits                 *int                   `json:"BusWidthBits"`                 // Bus Width in bits.
	AllowedSpeedsMHz             *int                   `json:"AllowedSpeedsMHz"`             // Speed bins supported by this Memory.
	FirmwareRevision             *string                `json:"FirmwareRevision"`             // Revision of firmware on the Memory controller.
	FirmwareAPIVersion           *string                `json:"FirmwareAPIVersion"`           // Version of API supported by the firmware.
	FunctionClasses              *[]string              `json:"FunctionClasses"`              // Function Classes by the Memory.
	VendorID                     *string                `json:"VendorID"`                     // Vendor ID.
	DeviceID                     *string                `json:"DeviceID"`                     // Device ID.
	SubsystemVendorID            *string                `json:"SubsystemVendorID"`            // SubSystem Vendor ID.
	SubsystemDeviceID            *string                `json:"SubsystemDeviceID"`            // Subsystem Device ID.
	MaxTDPMilliWatts             *[]int                 `json:"MaxTDPMilliWatts"`             // aximum TDPs in milli Watts.
	SecurityCapabilities         *SecurityCapabilities  `json:"SecurityCapabilities"`         // This object contains security capabilities of the Memory.
	SpareDeviceCount             *int                   `json:"SpareDeviceCount"`             // Number of unused spare devices available in the Memory.
	RankCount                    *int                   `json:"RankCount"`                    // Number of ranks available in the Memory.
	DeviceLocator                *string                `json:"DeviceLocator"`                // Location of the Memory in the platform.
	MemoryLocation               *MemoryLocation        `json:"MemoryLocation"`               // Memory connection information to sockets and memory controllers.
	ErrorCorrection              *string                `json:"ErrorCorrection"`              // Error correction scheme supported for this memory.
	OperatingSpeedMhz            *int                   `json:"OperatingSpeedMhz"`            // Operating speed of Memory in MHz.
	VolatileRegionSizeLimitMiB   *int                   `json:"VolatileRegionSizeLimitMiB"`   // Total size of volatile regions in MiB.
	PersistentRegionSizeLimitMiB *int                   `json:"PersistentRegionSizeLimitMiB"` // Total size of persistent regions in MiB.
	Regions                      *RegionSet             `json:"Regions"`                      // Memory regions information within the Memory.
	OperatingMemoryModes         *string                `json:"OperatingMemoryModes"`         // Memory modes supported by the Memory.
	PowerManagementPolicy        *PowerManagementPolicy `json:"PowerManagementPolicy"`        // Power management policy information.
	IsSpareDeviceEnabled         *bool                  `json:"IsSpareDeviceEnabled"`         // Spare device enabled status.
	IsRankSpareEnabled           *bool                  `json:"IsRankSpareEnabled"`           // Rank spare enabled status.
	Metrics                      *MemoryMetrics         `json:"Metrics"`                      // A reference to the Metrics associated with this Memory.
	VolatileRegionNumberLimit    *int                   `json:"VolatileRegionNumberLimit"`    // Total number of volatile regions this Memory can support.
	PersistentRegionNumberLimit  *int                   `json:"PersistentRegionNumberLimit"`  // Total number of persistent regions this Memory can support.
	VolatileRegionSizeMaxMiB     *int                   `json:"VolatileRegionSizeMaxMiB"`     // Maximum size of a single volatile region in MiB.
	PersistentRegionSizeMaxMiB   *int                   `json:"PersistentRegionSizeMaxMiB"`   // Maximum size of a single persistent region in MiB.
	AllocationIncrementMiB       *int                   `json:"AllocationIncrementMiB"`       // The size of the smallest unit of allocation for a memory region, thus it is the multiple in which regions are actually reserved.
	AllocationAlignmentMiB       *int                   `json:"AllocationAlignmentMiB"`       // The boundary which memory regions are allocated on, measured in MiB.
}
