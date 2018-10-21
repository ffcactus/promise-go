package dto

import (
	"promise/base"
)

// AllocateIPv4Request is the DTO to allocate an IP.
type AllocateIPv4Request struct {
	Key *string `json:"Key"`
}

// NewInstance returns a new instance.
func (AllocateIPv4Request) NewInstance() base.RequestInterface {
	return new(AllocateIPv4Request)
}

// IsValid return if the request is valid.
func (dto *AllocateIPv4Request) IsValid() *base.ErrorResponse {
	return nil
}

// String return the name for debug.
func (dto AllocateIPv4Request) String() string {
	if dto.Key == nil {
		return ""
	}
	return *dto.Key
}
