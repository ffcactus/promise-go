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

// Thermal Thermal object.
type Thermal struct {
	ServerRef string
	EmbeddedResource
	Fans []Fan `gorm:"ForeignKey:ThermalRef"` // This is the definition for fans.
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

// Load will load data from model.
func (e *Thermal) Load(m *model.Thermal) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)

	fans := []Fan{}
	for i := range m.Fans {
		e := Fan{}
		m := m.Fans[i]
		updateMemberEntity(&e.Member, &m.Member)
		updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
		updateThresholdEntity(&e.Threshold, &m.Threshold)
		e.Reading = m.Reading
		e.MinReadingRange = m.MinReadingRange
		e.MaxReadingRange = m.MaxReadingRange
		e.ReadingUnits = m.ReadingUnits
		fans = append(fans, e)
	}
	e.Fans = fans
}
