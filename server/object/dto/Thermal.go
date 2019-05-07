package dto

import (
	"promise/server/object/model"
)

// Fan This is the definition for fans.
type Fan struct {
	MemberResponse
	ProductInfoResponse
	ThresholdResponse
	Reading         *int    `json:"Reading"`         // Current fan speed.
	MinReadingRange *int    `json:"MinReadingRange"` // Minimum value for Reading.
	MaxReadingRange *int    `json:"MaxReadingRange"` // Maximum value for Reading.
	ReadingUnits    *string `json:"ReadingUnits"`    // Units in which the reading and thresholds are measured.
}

// Thermal is the DTO.
type Thermal struct {
	ResourceResponse
	Fans []Fan `json:"Fans"` // This is the definition for fans.
}

// Load will load data from model.
func (dto *Thermal) Load(m *model.Thermal) {
	dto.LoadResourceResponse(&m.Resource)
	dto.Fans = make([]Fan, 0)
	for i := range m.Fans {
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
