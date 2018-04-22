package entity

import (
	"promise/server/object/model"
)

// Voltage This is the definition for voltage sensors.
type Voltage struct {
	PowerRef uint
	EmbeddedResource
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
	PowerRef uint
	EmbeddedResource
	ProductInfo
	PowerConsumedWatts               *float64 // The actual power being consumed by the chassis.
	PowerRequestedWatts              *float64 // The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use.
	PowerAvailableWatts              *float64 // The amount of power not already budgeted and therefore available for additional allocation. (powerCapacity - powerAllocated).  This indicates how much reserve power capacity is left.
	PowerCapacityWatts               *float64 // The total amount of power available to the chassis for allocation. This may the power supply capacity, or power budget assigned to the chassis from an up-stream chassis.
	PowerAllocatedWatts              *float64 // The total amount of power that has been allocated (or budegeted)to  chassis resources.
	PowerMetricsMinConsumedWatts     *float64
	PowerMetricsMaxConsumedWatts     *float64
	PowerMetricsAverageConsumedWatts *float64
	PowerLimitLimitInWatts           *float64
	PowerLimitLimitException         *string
	PowerLimitCorrectionInMs         *float64
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

// PowerSupply Power supply object.
type PowerSupply struct {
	PowerRef uint
	EmbeddedResource
	ProductInfo
	PowerSupplyType      *string  // The Power Supply type (AC or DC).
	LineInputVoltageType *string  // The line voltage type supported as an input to this Power Supply.
	LineInputVoltage     *float64 // The line input voltage at which the Power Supply is operating.
	PowerCapacityWatts   *float64 // The maximum capacity of this Power Supply.
	LastPowerOutputWatts *float64 // The average power output of this Power Supply.
	FirmwareVersion      *string  // The firmware version for this Power Supply.
	IndicatorLed         *string  // The state of the indicator LED, used to identify the power supply.
}

// Power Power object.
type Power struct {
	ServerRef string
	EmbeddedResource
	PowerControl  []PowerControl `gorm:"ForeignKey:PowerRef"` // This is the definition for power control function (power reading/limiting).
	Voltages      []Voltage      `gorm:"ForeignKey:PowerRef"` // This is the definition for voltage sensors.
	PowerSupplies []PowerSupply  `gorm:"ForeignKey:PowerRef"` // Details of the power supplies associated with this system or device.
	Redundancy    []Redundancy   `gorm:"ForeignKey:Ref"`      // Redundancy information for the power subsystem of this system or device.
}

// ToModel will create a new model from entity.
func (e *PowerControl) ToModel() *model.PowerControl {
	m := model.PowerControl{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.PowerConsumedWatts = e.PowerConsumedWatts
	m.PowerRequestedWatts = e.PowerRequestedWatts
	m.PowerAvailableWatts = e.PowerAvailableWatts
	m.PowerCapacityWatts = e.PowerCapacityWatts
	m.PowerAllocatedWatts = e.PowerAllocatedWatts

	m.PowerMetrics = new(model.PowerMetrics)
	m.PowerMetrics.MinConsumedWatts = e.PowerMetricsMinConsumedWatts
	m.PowerMetrics.MaxConsumedWatts = e.PowerMetricsMaxConsumedWatts
	m.PowerMetrics.AverageConsumedWatts = e.PowerMetricsAverageConsumedWatts

	m.PowerLimit = new(model.PowerLimit)
	m.PowerLimit.LimitInWatts = e.PowerLimitLimitInWatts
	m.PowerLimit.LimitException = e.PowerLimitLimitException
	m.PowerLimit.CorrectionInMs = e.PowerLimitCorrectionInMs
	return &m
}

// ToModel will create a new model from entity.
func (e *Voltage) ToModel() *model.Voltage {
	m := model.Voltage{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.SensorNumber = e.SensorNumber
	m.ReadingVolts = e.ReadingVolts
	m.MinReadingRange = e.MinReadingRange
	m.MaxReadingRange = e.MaxReadingRange
	m.PhysicalContext = e.PhysicalContext
	return &m
}

// ToModel will create a new model from entity.
func (e *PowerSupply) ToModel() *model.PowerSupply {
	m := model.PowerSupply{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.PowerSupplyType = e.PowerSupplyType
	m.LineInputVoltageType = e.LineInputVoltageType
	m.LineInputVoltage = e.LineInputVoltage
	m.PowerCapacityWatts = e.PowerCapacityWatts
	m.LastPowerOutputWatts = e.LastPowerOutputWatts
	m.FirmwareVersion = e.FirmwareVersion
	m.IndicatorLed = e.IndicatorLed
	return &m
}