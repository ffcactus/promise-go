package dto

// IPv4Range is a IPv4 range.
type IPv4Range struct {
	Start string `json:"Start"`
	End   string `json:"End"`
	Total		*uint32		`json:"Total,omitempty"`
	Free		*uint32		`json:"Free,omitempty"`
	Allocatable *uint32		`json:"Allocatable,omitempty"`	
}

// IPv4PoolResource is the resource DTO.
type IPv4PoolResource struct {
	Name        string      `json:"Name"`
	Description *string     `json:"Description,omitempty"`
	Ranges      []IPv4Range `json:"Ranges"`
	SubnetMask  string      `json:"SubnetMask"`
	Gateway     string      `json:"Gateway"`
	Domain      string      `json:"Domain"`
	DNSServers  []string    `json:"DNSServers"`
	Total		*uint32		`json:"Total,omitempty"`
	Free		*uint32		`json:"Free,omitempty"`
	Allocatable *uint32		`json:"Allocatable,omitempty"`
}
