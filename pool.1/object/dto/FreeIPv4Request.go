package dto

import (
	"net"
	commonConstError "promise/common/object/consterror"
	commonDTO "promise/common/object/dto"
)

// FreeIPv4Request is the DTO to free an IP.
type FreeIPv4Request struct {
	commonDTO.PromiseRequest
	Address string `json:"Address"`
}

// Validate the request.
func (dto *FreeIPv4Request) Validate() error {
	if net.ParseIP(dto.Address) == nil {
		return commonConstError.ErrorDataConvert
	}
	return nil
}
