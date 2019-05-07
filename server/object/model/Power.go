package model

import (
	"fmt"
)

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
	IndicatorLed         *string  // The state of the indicator LED, used to identify the power supply.
}

// Power The power.
type Power struct {
	Resource
	PowerControl  []PowerControl // This is the definition for power control function (power reading/limiting).
	PowerSupplies []PowerSupply  // Details of the power supplies associated with this system or device.
	Redundancy    []Redundancy   // Redundancy information for the power subsystem of this system or device.
}

// String output debug info.
func (m Power) String() string {
	return fmt.Sprintf("control %d,supply %d", len(m.PowerControl), len(m.PowerSupplies))
}
