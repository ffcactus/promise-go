package dto

// AllocateIPv4Response is the DTO to allocate an IP.
type AllocateIPv4Response struct {
	Address string
	Pool    GetIPv4PoolResponse
}

// String return the debug info.
func (dto AllocateIPv4Response) String() string {
	return dto.Address
}
