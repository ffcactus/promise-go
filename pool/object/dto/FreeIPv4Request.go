package dto

import (
	"net"
	commonDTO "promise/common/object/dto"
)

// FreeIPv4Request is the DTO to free an IP.
type FreeIPv4Request struct {
	commonDTO.PromiseRequest
	Address string `json:"Address"`
}

// IsValid check if the request is valid.
func (dto *FreeIPv4Request) IsValid() bool {
	if net.ParseIP(dto.Address) == nil {
		return false
	}
	return true
}
