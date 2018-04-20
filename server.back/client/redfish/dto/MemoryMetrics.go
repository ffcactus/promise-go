package dto

// Alarm trip information about the memory.
type AlarmTrips struct {
	Temperature           *bool // Temperature threshold crossing alarm trip detected status.
	SpareBlock            *bool // Spare block capacity crossing alarm trip detected status.
	UncorrectableECCError *bool // Uncorrectable data error threshold crossing alarm trip detected status.
	CorrectableECCError   *bool // Correctable data error threshold crossing alarm trip detected status.
	AddressParityError    *bool // Address parity error detected status.
}

// This object contains the Memory metrics since last reset or ClearCurrentPeriod action.
type CurrentPeriod struct {
	BlocksRead    *int // Number of blocks read since reset.
	BlocksWritten *int // Number of blocks written since reset.
}

// This object contains the Memory metrics for the lifetime of the Memory.
type LifeTime struct {
	BlocksRead    *int // Number of blocks read for the lifetime of the Memory.
	BlocksWritten *int // Number of blocks written for the lifetime of the Memory.
}

// This type describes the health information of the memory.
type HealthData struct {
	RemainingSpareBlockPercentage *int        // Remaining spare blocks in percentage.
	LastShutdownSuccess           *bool       // Status of last shutdown.
	DataLossDetected              *bool       // Data loss detection status.
	PerformanceDegraded           *bool       // Performance degraded mode status.
	AlarmTrips                    *AlarmTrips // Alarm trip information about the memory.
	PredictedMediaLifeLeftPercent *int        // The percentage of reads and writes that are predicted to still be available for the media.
}

// MemoryMetrics contains usage and health statistics for a single Memory module or device instance.
type MemoryMetrics struct {
	Id             *string
	PageURI        *string
	State          *string
	Health         *string
	Name           *string
	Description    *string
	BlockSizeBytes *int           // Block size in bytes.
	CurrentPeriod  *CurrentPeriod // This object contains the Memory metrics since last reset or ClearCurrentPeriod action.
	LifeTime       *LifeTime      // This object contains the Memory metrics for the lifetime of the Memory.
	HealthData     *HealthData    // This object describes the health information of the memory.
}
