package dto

import (
	"promise/server/object/model"
)

// This is the definition for fans.
type Fan struct {
	MemberResponse
	ProductInfoResponse
	ThresholdResponse
	Reading         *int `json:"Reading"`         // Current fan speed.
	MinReadingRange *int `json:"MinReadingRange"` // Minimum value for Reading.
	MaxReadingRange *int `json:"MaxReadingRange"` // Maximum value for Reading.
	//	Redundancy      *Redundancy `json:"Redundancy"` // This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	ReadingUnits *string `json:"ReadingUnits"` // Units in which the reading and thresholds are measured.
}

// This is the definition for temperature sensors.
type Temperature struct {
	MemberResponse
	ThresholdResponse
	SensorNumber        *int `json:"SensorNumber"`        // A numerical identifier to represent the temperature sensor.
	ReadingCelsius      *int `json:"ReadingCelsius"`      // Temperature.
	MinReadingRangeTemp *int `json:"MinReadingRangeTemp"` // Minimum value for ReadingCelsius.
	MaxReadingRangeTemp *int `json:"MaxReadingRangeTemp"` // Maximum value for ReadingCelsius.
}

type Thermal struct {
	ResourceResponse
	Temperatures []Temperature `json:"Temperatures"` // This is the definition for temperature sensors.
	Fans         []Fan         `json:"Fans"`         // This is the definition for fans.
}

func (dto *Thermal) Load(m *model.Thermal) {
	dto.LoadResourceResponse(&m.Resource)
	dto.Temperatures = make([]Temperature, 0)
	dto.Fans = make([]Fan, 0)
	for i, _ := range m.Temperatures {
		d := Temperature{}
		m := m.Temperatures[i]
		d.LoadMemberResponse(&m.Member)
		d.LoadThresholdResponse(&m.Threshold)
		d.SensorNumber = m.SensorNumber
		d.ReadingCelsius = m.ReadingCelsius
		d.MinReadingRangeTemp = m.MinReadingRangeTemp
		d.MaxReadingRangeTemp = m.MaxReadingRangeTemp
		dto.Temperatures = append(dto.Temperatures, d)
	}
	for i, _ := range m.Temperatures {
		d := Temperature{}
		m := m.Temperatures[i]
		d.LoadMemberResponse(&m.Member)
		d.LoadThresholdResponse(&m.Threshold)
		d.SensorNumber = m.SensorNumber
		d.ReadingCelsius = m.ReadingCelsius
		d.MinReadingRangeTemp = m.MinReadingRangeTemp
		d.MaxReadingRangeTemp = m.MaxReadingRangeTemp
		dto.Temperatures = append(dto.Temperatures, d)
	}
	for i, _ := range m.Fans {
		d := Fan{}
		m := m.Fans[i]
		d.LoadMemberResponse(&m.Member)
		d.LoadProductInfoResponse(&m.ProductInfo)
		d.LoadThresholdResponse(&m.Threshold)
		d.Reading = m.Reading
		d.MinReadingRange = m.MinReadingRange
		d.MaxReadingRange = m.MaxReadingRange
		d.ReadingUnits = m.ReadingUnits
		dto.Fans = append(dto.Fans, d)
	}
}
