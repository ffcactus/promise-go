package dto

// Fan is fan hardware.
type Fan struct {
	Member
	ProductInfo
	Threshold
	Reading         *int    // Current fan speed.
	MinReadingRange *int    // Minimum value for Reading.
	MaxReadingRange *int    // Maximum value for Reading.
	ReadingUnits    *string // Units in which the reading and thresholds are measured.
}

// GetThermalResponse is the thermal system response from redfish interface.
type GetThermalResponse struct {
	Resource
	Fans []Fan // This is the definition for fans.
}
