package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/model"
)

// AllocateIPv4Request is the DTO to allocate an IP.
type AllocateIPv4Request struct {
	base.ActionRequest
	Key *string `json:"Key"`
}

// IsValid return if the request is valid.
func (dto *AllocateIPv4Request) IsValid() *base.Message {
	return nil
}

// GetDebugName return the name for debug.
func (dto *AllocateIPv4Request) GetDebugName() string {
	if dto.Key == nil {
		return ""
	}
	return *dto.Key
}

// UpdateModel Update the model.
func (dto *AllocateIPv4Request) UpdateModel(i base.ModelInterface) error {
	_, ok := i.(*model.IPv4Pool)
	if !ok {
		log.Error("AllocateIPv4Request.UpdateModel() convert interface failed.")
		return base.ErrorDataConvert
	}

	return nil
}
