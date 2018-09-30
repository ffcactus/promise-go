package mm920

// GetChassisEnclosureResponse is DTO.
// Represents response from get /redfish/v1/Chassis/Enclosure
type GetChassisEnclosureResponse struct {
	SerialNumber string
	PartNumber   string
}

// GetRedfishV1Response is DTO.
// Represents response from get /redfish/v1
type GetRedfishV1Response struct {
	UUID string
}

// GetBladeChassisResponse is DTO.
// Represents response from get /redfish/v1/chassis/blade{:id}
type GetBladeChassisResponse struct {
	Status struct {
		State string
	}
	Model        string
	SerialNumber string
	PartNumber   string
	Oem          struct {
		Huawei struct {
			Width  int
			Height int
		}
	}
}
