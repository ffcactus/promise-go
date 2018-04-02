package dto

// IPv4Range is a IPv4 range.
type IPv4Range struct {
	Start string `json:"Start"`
	End   string `json:"End"`
}

// IPv4PoolResource is the resource DTO.
type IPv4PoolResource struct {
	Name        string      `json:"Name"`
	Description *string     `json:"Description"`
	Ranges      []IPv4Range `json:"Ranges"`
	SubnetMask  string      `json:"SubnetMask"`
	Gateway     string      `json:"Gateway"`
	Domain      string      `json:"Domain"`
	DNSServers  []string    `json:"DNSServers"`
}
