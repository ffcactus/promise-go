package dto

// IPv4RangeRequest is a IPv4 range in request.
type IPv4RangeRequest struct {
	Start       string  `json:"Start"`
	End         string  `json:"End"`
	Total       *uint32 `json:"Total"`
	Free        *uint32 `json:"Free"`
	Allocatable *uint32 `json:"Allocatable"`
}

// IPv4RangeResponse is a IPv4 range in response.
type IPv4RangeResponse struct {
	Start       string `json:"Start"`
	End         string `json:"End"`
	Total       uint32 `json:"Total"`
	Free        uint32 `json:"Free"`
	Allocatable uint32 `json:"Allocatable"`
}

// IPv4PoolResource is the resource DTO.
type IPv4PoolResource struct {
	Name        string    `json:"Name"`
	Description *string   `json:"Description,omitempty"`
	SubnetMask  *string   `json:"SubnetMask,omitempty"`
	Gateway     *string   `json:"Gateway,omitempty"`
	Domain      *string   `json:"Domain,omitempty"`
	DNSServers  *[]string `json:"DNSServers,omitempty"`
}
