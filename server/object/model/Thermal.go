package model

import (
	"fmt"
)

// Fan This is the definition for fans.
type Fan struct {
	Member
	ProductInfo
	Threshold
	Reading         *int    // Current fan speed.
	MinReadingRange *int    // Minimum value for Reading.
	MaxReadingRange *int    // Maximum value for Reading.
	ReadingUnits    *string // Units in which the reading and thresholds are measured.
}

// Thermal Thermal object.
type Thermal struct {
	Resource
	Fans []Fan // This is the definition for fans.
}

// String output debug info.
func (m Thermal) String() string {
	return fmt.Sprintf("fan %d", len(m.Fans))
}
