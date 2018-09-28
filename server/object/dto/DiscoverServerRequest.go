package dto

import (
	"promise/base"
)

// DiscoverServerRequest The request body of POST server.
type DiscoverServerRequest struct {
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Hostname    string  `json:"Hostname"`
	Username    string  `json:"Username"`
	Password    string  `json:"Password"`
}

// NewInstance returns a new instance.
func (dto *DiscoverServerRequest) NewInstance() base.RequestInterface {
	return new(DiscoverServerRequest)
}

// IsValid return if the request is valid.
func (dto *DiscoverServerRequest) IsValid() *base.ErrorResponse {
	if dto.Hostname == "" || dto.Username == "" || dto.Password == "" {
		return base.NewErrorResponseInvalidRequest()
	}
	return nil
}

// String return the name for debug.
func (dto DiscoverServerRequest) String() string {
	return dto.Hostname
}
