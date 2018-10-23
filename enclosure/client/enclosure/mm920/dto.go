package mm920

// Collection represents the collection response.
type Collection struct {
	Count   int `json:"Members@odata.count"`
	Members []struct {
		ID string `json:"@odata.id"`
	}
	NextLink string `json:"Members@odata.nextLink"`
}

// GetChassisEnclosureResponse is DTO.
// It represents response from get /redfish/v1/Chassis/Enclosure
type GetChassisEnclosureResponse struct {
	SerialNumber string
	PartNumber   string
}

// GetRedfishV1Response is DTO.
// It represents response from get /redfish/v1
type GetRedfishV1Response struct {
	UUID string
}

// GetBladeChassisResponse is DTO.
// It represents response from get /redfish/v1/chassis/blade{:id}
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

// GetSwiChassisResponse is DTO,.
// It represents resposne from get /redfish/v1/chassis/swi{id}
type GetSwiChassisResponse struct {
	Status struct {
		State string
	}
	Model string
}

// GetSwiSystemResponse is DTO.
// It represents resposne from get /redfish/v1/systems/swi{id}
type GetSwiSystemResponse struct {
	SerialNumber string
}
