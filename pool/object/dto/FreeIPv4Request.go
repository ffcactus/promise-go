package dto

import (
	log "github.com/sirupsen/logrus"
	"net"
	"promise/base"
	"promise/pool/object/errorResp"
	"promise/pool/object/model"
)

// FreeIPv4Request is the DTO to free an IP.
type FreeIPv4Request struct {
	Address string `json:"Address"`
}

// NewInstance returns a new instance.
func (FreeIPv4Request) NewInstance() base.RequestInterface {
	return new(FreeIPv4Request)
}

// IsValid return if the request is valid.
func (dto *FreeIPv4Request) IsValid() *base.ErrorResponse {
	if net.ParseIP(dto.Address) == nil {
		return errorResp.NewErrorResponseIPv4FormatError()
	}
	return nil
}

// String return the name for debug.
func (dto FreeIPv4Request) String() string {
	return dto.Address
}

// UpdateModel Update the model.
func (dto *FreeIPv4Request) UpdateModel(i base.ModelInterface) error {
	_, ok := i.(*model.IPv4Pool)
	if !ok {
		log.Error("FreeIPv4Request.UpdateModel() convert interface failed.")
		return base.ErrorDataConvert
	}

	return nil
}
