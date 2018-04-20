package dto

import (
	"promise/server/object/model"
)

// Voltage This is the definition for voltage sensors.
type Voltage struct {
	ResourceResponse
	ThresholdResponse
	SensorNumber    *int     `json:"SensorNumber,omitempty"`    // A numerical identifier to represent the voltage sensor.
	ReadingVolts    *float64 `json:"ReadingVolts,omitempty"`    // The present reading of the voltage sensor.
	MinReadingRange *float64 `json:"MinReadingRange,omitempty"` // Minimum value for this Voltage sensor.
	MaxReadingRange *float64 `json:"MaxReadingRange,omitempty"` // Maximum value for this Voltage sensor.
	PhysicalContext *string  `json:"PhysicalContext,omitempty"`
}

// PowerMetrics Power readings for this chassis.
type PowerMetrics struct {
	MinConsumedWatts     *float64 `json:"MinConsumedWatts,omitempty"`     // The lowest power consumption level over the measurement window (the last IntervalInMin minutes).
	MaxConsumedWatts     *float64 `json:"MaxConsumedWatts,omitempty"`     // The highest power consumption level that has occured over the measurement window (the last IntervalInMin minutes).
	AverageConsumedWatts *float64 `json:"AverageConsumedWatts,omitempty"` // The average power level over the measurement window (the last IntervalInMin minutes).
}

// PowerLimit This object contains power limit status and configuration information for the chassis.
type PowerLimit struct {
	LimitInWatts   *float64 `json:"LimitInWatts"`             // The Power limit in watts. Set to null to disable power capping.
	LimitException *string  `json:"LimitException"`           // The action that is taken if the power cannot be maintained below the LimitInWatts.
	CorrectionInMs *float64 `json:"CorrectionInMs,omitempty"` // The time required for the limiting process to reduce power consumption to below the limit.
}

// PowerControl This is the definition for power control function (power reading/limiting).
type PowerControl struct {
	ResourceResponse
	ProductInfoResponse
	PowerConsumedWatts  *float64      `json:"PowerConsumedWatts,omitempty"`  // The actual power being consumed by the chassis.
	PowerRequestedWatts *float64      `json:"PowerRequestedWatts,omitempty"` // The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use.
	PowerAvailableWatts *float64      `json:"PowerAvailableWatts,omitempty"` // The amount of power not already budgeted and therefore available for additional allocation. (powerCapacity - powerAllocated).  This indicates how much reserve power capacity is left.
	PowerCapacityWatts  *float64      `json:"PowerCapacityWatts,omitempty"`  // The total amount of power available to the chassis for allocation. This may the power supply capacity, or power budget assigned to the chassis from an up-stream chassis.
	PowerAllocatedWatts *float64      `json:"PowerAllocatedWatts,omitempty"` // The total amount of power that has been allocated (or budegeted)to  chassis resources.
	PowerMetrics        *PowerMetrics `json:"PowerMetrics,omitempty"`        // Power readings for this chassis.
	PowerLimit          *PowerLimit   `json:"PowerLimit,omitempty"`          // The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use.
	// RelatedItem         *[]string
}

// InputRange This type shall describe an input range that the associated power supply is able to utilize.
type InputRange struct {
	InputType          *string `json:"InputType,omitempty"`          // The Input type (AC or DC).
	MinimumVoltage     *int    `json:"MinimumVoltage,omitempty"`     // The minimum line input voltage at which this power supply input range is effective.
	MinimumFrequencyHz *int    `json:"MinimumFrequencyHz,omitempty"` // The minimum line input frequency at which this power supply input range is effective.
	MaximumFrequencyHz *int    `json:"MaximumFrequencyHz,omitempty"` // The maximum line input frequency at which this power supply input range is effective.
	OutputWattage      *int    `json:"OutputWattage,omitempty"`      // The maximum capacity of this Power Supply when operating in this input range.
}

// PowerSupply is DTO.
type PowerSupply struct {
	ResourceResponse
	ProductInfoResponse
	PowerSupplyType      *string  `json:"PowerSupplyType,omitempty"`      // The Power Supply type (AC or DC).
	LineInputVoltageType *string  `json:"LineInputVoltageType,omitempty"` // The line voltage type supported as an input to this Power Supply.
	LineInputVoltage     *float64 `json:"LineInputVoltage,omitempty"`     // The line input voltage at which the Power Supply is operating.
	PowerCapacityWatts   *float64 `json:"PowerCapacityWatts,omitempty"`   // The maximum capacity of this Power Supply.
	LastPowerOutputWatts *float64 `json:"LastPowerOutputWatts,omitempty"` // The average power output of this Power Supply.
	FirmwareVersion      *string  `json:"FirmwareVersion,omitempty"`      // The firmware version for this Power Supply.
	//	RelatedItem          *[]string     // The ID(s) of the resources associated with this Power Limit.
	// Redundancy   *Redundancy   // This structure is used to show redundancy for power supplies.  The Component ids will reference the members of the redundancy groups.
	// InputRange   *[]InputRange // This is the input ranges that the power supply can use.
	IndicatorLed *string `json:"IndicatorLed,omitempty"` // The state of the indicator LED, used to identify the power supply.
}

// Power is DTO.
type Power struct {
	ResourceResponse
	PowerControl  []PowerControl `json:"PowerControl"`  // This is the definition for power control function (power reading/limiting).
	Voltages      []Voltage      `json:"Voltages"`      // This is the definition for voltage sensors.
	PowerSupplies []PowerSupply  `json:"PowerSupplies"` // Details of the power supplies associated with this system or device.
	Redundancy    []Redundancy   `json:"Redundancy"`    // Redundancy information for the power subsystem of this system or device.
}

// Load will load data from model.
func (dto *Power) Load(m *model.Power) {
	dto.LoadResourceResponse(&m.Resource)
	dto.PowerControl = make([]PowerControl, 0)
	dto.Voltages = make([]Voltage, 0)
	dto.PowerSupplies = make([]PowerSupply, 0)
	dto.Redundancy = make([]Redundancy, 0)
	// PowerControl
	if m.PowerControl != nil {
		for i := range *m.PowerControl {
			each := PowerControl{}
			powerControl := (*m.PowerControl)[i]
			each.LoadResourceResponse(&powerControl.Resource)
			each.LoadProductInfoResponse(&powerControl.ProductInfo)
			each.PowerConsumedWatts = powerControl.PowerConsumedWatts
			each.PowerRequestedWatts = powerControl.PowerRequestedWatts
			each.PowerAvailableWatts = powerControl.PowerAvailableWatts
			each.PowerCapacityWatts = powerControl.PowerCapacityWatts
			each.PowerAllocatedWatts = powerControl.PowerAllocatedWatts
			if powerControl.PowerMetrics != nil {
				each.PowerMetrics = new(PowerMetrics)
				each.PowerMetrics.MinConsumedWatts = powerControl.PowerMetrics.MinConsumedWatts
				each.PowerMetrics.MaxConsumedWatts = powerControl.PowerMetrics.MaxConsumedWatts
				each.PowerMetrics.AverageConsumedWatts = powerControl.PowerMetrics.AverageConsumedWatts
			}
			if powerControl.PowerLimit != nil {
				each.PowerLimit = new(PowerLimit)
				each.PowerLimit.LimitInWatts = powerControl.PowerLimit.LimitInWatts
				each.PowerLimit.LimitException = powerControl.PowerLimit.LimitException
				each.PowerLimit.CorrectionInMs = powerControl.PowerLimit.CorrectionInMs
			}
			dto.PowerControl = append(dto.PowerControl, each)
		}
	}
	// Voltages
	if m.Voltages != nil {
		for i := range *m.Voltages {
			each := Voltage{}
			voltage := (*m.Voltages)[i]
			each.LoadResourceResponse(&voltage.Resource)
			each.LoadThresholdResponse(&voltage.Threshold)
			each.SensorNumber = voltage.SensorNumber
			each.ReadingVolts = voltage.ReadingVolts
			each.MinReadingRange = voltage.MinReadingRange
			each.MaxReadingRange = voltage.MaxReadingRange
			each.PhysicalContext = voltage.PhysicalContext
			dto.Voltages = append(dto.Voltages, each)
		}
	}
	// PowerSupply
	if m.PowerSupplies != nil {
		for i := range *m.PowerSupplies {
			each := PowerSupply{}
			powerSupplies := (*m.PowerSupplies)[i]
			each.LoadResourceResponse(&powerSupplies.Resource)
			each.LoadProductInfoResponse(&powerSupplies.ProductInfo)
			each.PowerSupplyType = powerSupplies.PowerSupplyType
			each.LineInputVoltageType = powerSupplies.LineInputVoltageType
			each.LineInputVoltage = powerSupplies.LineInputVoltage
			each.PowerCapacityWatts = powerSupplies.PowerCapacityWatts
			each.LastPowerOutputWatts = powerSupplies.LastPowerOutputWatts
			each.FirmwareVersion = powerSupplies.FirmwareVersion
			each.IndicatorLed = powerSupplies.IndicatorLed
			dto.PowerSupplies = append(dto.PowerSupplies, each)
		}
	}
	// Redundancy
	if m.Redundancy != nil {
		for i := range *m.Redundancy {
			each := Redundancy{}
			redundancy := (*m.Redundancy)[i]
			each.Load(&redundancy)
			dto.Redundancy = append(dto.Redundancy, each)
		}
	}
}
