package dto

import (
	"promise/base"
)

// DiscoverEnclosureRequest is the DTO.
type DiscoverEnclosureRequest struct {
	Address  string
	Username string
	Password string
	Force    bool
}

// NewInstance returns a new instance.
func (dto *DiscoverEnclosureRequest) NewInstance() base.RequestInterface {
	return new(DiscoverEnclosureRequest)
}

// IsValid return if the request is valid.
func (dto *DiscoverEnclosureRequest) IsValid() *base.ErrorResponse {
	if dto.Address == "" || dto.Username == "" || dto.Password == "" {
		return base.NewErrorResponseInvalidRequest()
	}
	return nil
}

// String return the name for debug.
func (dto DiscoverEnclosureRequest) String() string {
	return dto.Address
}
