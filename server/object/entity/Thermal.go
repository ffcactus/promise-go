package entity

import (
	"promise/server/object/model"
)

// Fan This is the definition for fans.
type Fan struct {
	ThermalRef uint
	Member
	ProductInfo
	Threshold
	Reading         *int // Current fan speed.
	MinReadingRange *int // Minimum value for Reading.
	MaxReadingRange *int // Maximum value for Reading.
	// Redundancy      *Redundancy // This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	ReadingUnits *string // Units in which the reading and thresholds are measured.
}

// Temperature This is the definition for temperature sensors.
type Temperature struct {
	ThermalRef uint
	Member
	Threshold
	SensorNumber        *int // A numerical identifier to represent the temperature sensor.
	ReadingCelsius      *int // Temperature.
	MinReadingRangeTemp *int // Minimum value for ReadingCelsius.
	MaxReadingRangeTemp *int // Maximum value for ReadingCelsius.
}

// Thermal Thermal object.
type Thermal struct {
	ServerRef string
	EmbeddedResource
	Temperatures []Temperature `gorm:"ForeignKey:ThermalRef"` // This is the definition for temperature sensors.
	Fans         []Fan         `gorm:"ForeignKey:ThermalRef"` // This is the definition for fans.
}

// ToModel will create a new model from entity.
func (e *Temperature) ToModel() *model.Temperature {
	m := new(model.Temperature)
	createMemberModel(&e.Member, &m.Member)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.SensorNumber = e.SensorNumber
	m.ReadingCelsius = e.ReadingCelsius
	m.MinReadingRangeTemp = e.MinReadingRangeTemp
	m.MaxReadingRangeTemp = e.MaxReadingRangeTemp
	return m
}

// ToModel will create a new model from entity.
func (e *Fan) ToModel() *model.Fan {
	m := new(model.Fan)
	createMemberModel(&e.Member, &m.Member)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.Reading = e.Reading
	m.MinReadingRange = e.MinReadingRange
	m.MaxReadingRange = e.MaxReadingRange
	m.ReadingUnits = e.ReadingUnits
	return m
}
