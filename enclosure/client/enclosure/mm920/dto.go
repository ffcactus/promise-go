package mm920

// GetChassisEnclosureResponse is DTO.
// Represents response from get /redfish/v1/Chassis/Enclosure
type GetChassisEnclosureResponse struct {
	SerialNumber string
	PartNumber string
}

// GetRedfishV1Response is DTO.
// Represents response from get /redfish/v1
type GetRedfishV1Response struct {
	UUID string
}