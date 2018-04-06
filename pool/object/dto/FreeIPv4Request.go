package dto

import (
	"net"
	commonDTO "promise/common/object/dto"
	commonMessage "promise/common/object/message"
	"promise/pool/object/message"
)

// FreeIPv4Request is the DTO to free an IP.
type FreeIPv4Request struct {
	commonDTO.PromiseRequest
	Address string `json:"Address"`
}

// Validate the request.
func (dto *FreeIPv4Request) Validate() *commonMessage.Message {
	if net.ParseIP(dto.Address) == nil {
		m := message.NewIPv4FormatError()
		return &m
	}
	return nil
}
