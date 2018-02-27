package model

// Voltage This is the definition for voltage sensors.
type Voltage struct {
	Resource
	Threshold
	SensorNumber    *int     // A numerical identifier to represent the voltage sensor.
	ReadingVolts    *float64 // The present reading of the voltage sensor.
	MinReadingRange *float64 // Minimum value for this Voltage sensor.
	MaxReadingRange *float64 // Maximum value for this Voltage sensor.
	PhysicalContext *string
	//	RelatedItem     *[]string // Describes the areas or devices to which this voltage measurement applies.
}

// PowerMetrics Power readings for this chassis.
type PowerMetrics struct {
	MinConsumedWatts     *float64 // The lowest power consumption level over the measurement window (the last IntervalInMin minutes).
	MaxConsumedWatts     *float64 // The highest power consumption level that has occured over the measurement window (the last IntervalInMin minutes).
	AverageConsumedWatts *float64 // The average power level over the measurement window (the last IntervalInMin minutes).
}

// PowerLimit This object contains power limit status and configuration information for the chassis.
type PowerLimit struct {
	LimitInWatts   *float64 // The Power limit in watts. Set to null to disable power capping.
	LimitException *string  // The action that is taken if the power cannot be maintained below the LimitInWatts.
	CorrectionInMs *float64 // The time required for the limiting process to reduce power consumption to below the limit.
}

// PowerControl This is the definition for power control function (power reading/limiting).
type PowerControl struct {
	Resource
	ProductInfo
	PowerConsumedWatts  *float64      // The actual power being consumed by the chassis.
	PowerRequestedWatts *float64      // The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use.
	PowerAvailableWatts *float64      // The amount of power not already budgeted and therefore available for additional allocation. (powerCapacity - powerAllocated).  This indicates how much reserve power capacity is left.
	PowerCapacityWatts  *float64      // The total amount of power available to the chassis for allocation. This may the power supply capacity, or power budget assigned to the chassis from an up-stream chassis.
	PowerAllocatedWatts *float64      // The total amount of power that has been allocated (or budegeted)to  chassis resources.
	PowerMetrics        *PowerMetrics // Power readings for this chassis.
	PowerLimit          *PowerLimit   // The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use.
	// RelatedItem         *[]string
}

// InputRange This type shall describe an input range that the associated power supply is able to utilize.
type InputRange struct {
	InputType          *string // The Input type (AC or DC).
	MinimumVoltage     *int    // The minimum line input voltage at which this power supply input range is effective.
	MinimumFrequencyHz *int    // The minimum line input frequency at which this power supply input range is effective.
	MaximumFrequencyHz *int    // The maximum line input frequency at which this power supply input range is effective.
	OutputWattage      *int    // The maximum capacity of this Power Supply when operating in this input range.
}

// PowerSupply The power supply
type PowerSupply struct {
	Resource
	ProductInfo
	PowerSupplyType      *string  // The Power Supply type (AC or DC).
	LineInputVoltageType *string  // The line voltage type supported as an input to this Power Supply.
	LineInputVoltage     *float64 // The line input voltage at which the Power Supply is operating.
	PowerCapacityWatts   *float64 // The maximum capacity of this Power Supply.
	LastPowerOutputWatts *float64 // The average power output of this Power Supply.
	FirmwareVersion      *string  // The firmware version for this Power Supply.
	// RelatedItem          *[]string     // The ID(s) of the resources associated with this Power Limit.
	// Redundancy   *Redundancy   // This structure is used to show redundancy for power supplies.  The Component ids will reference the members of the redundancy groups.
	// InputRange   *[]InputRange // This is the input ranges that the power supply can use.
	IndicatorLed *string // The state of the indicator LED, used to identify the power supply.
}

// Power The power.
type Power struct {
	Resource
	PowerControl  *[]PowerControl // This is the definition for power control function (power reading/limiting).
	Voltages      *[]Voltage      // This is the definition for voltage sensors.
	PowerSupplies *[]PowerSupply  // Details of the power supplies associated with this system or device.
	Redundancy    *[]Redundancy   // Redundancy information for the power subsystem of this system or device.
}
