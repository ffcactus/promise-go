package model

// Fan This is the definition for fans.
type Fan struct {
	Member
	ProductInfo
	Threshold
	Reading         *int        // Current fan speed.
	MinReadingRange *int        // Minimum value for Reading.
	MaxReadingRange *int        // Maximum value for Reading.
	Redundancy      *Redundancy // This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	ReadingUnits    *string     // Units in which the reading and thresholds are measured.
}

// Temperature This is the definition for temperature sensors.
type Temperature struct {
	Member
	Threshold
	SensorNumber        *int // A numerical identifier to represent the temperature sensor.
	ReadingCelsius      *int // Temperature.
	MinReadingRangeTemp *int // Minimum value for ReadingCelsius.
	MaxReadingRangeTemp *int // Maximum value for ReadingCelsius.
}

// Thermal Thermal object.
type Thermal struct {
	Resource
	Temperatures []Temperature // This is the definition for temperature sensors.
	Fans         []Fan         // This is the definition for fans.
}
