package dto

import (
	log "github.com/sirupsen/logrus"
	"net"
	"promise/base"
	"promise/pool/object/model"
)

// FreeIPv4Request is the DTO to free an IP.
type FreeIPv4Request struct {
	base.ActionRequest
	Address string `json:"Address"`
}

// IsValid return if the request is valid.
func (dto *FreeIPv4Request) IsValid() *base.Message {
	if net.ParseIP(dto.Address) == nil {
		message := base.NewMessageIPv4FormatError()
		return &message
	}
	return nil
}

// GetDebugName return the name for debug.
func (dto *FreeIPv4Request) GetDebugName() string {
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
