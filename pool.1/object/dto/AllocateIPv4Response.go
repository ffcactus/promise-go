package dto

// AllocateIPv4Response is the DTO to allocate an IP.
type AllocateIPv4Response struct {
	Address string
	Pool    GetIPv4PoolResponse
}
