package dto

import (
	"promise/base"
)

// DiscoverServerRequest The request body of POST server.
type DiscoverServerRequest struct {
	Hostname string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// NewInstance returns a new instance.
func (dto *DiscoverServerRequest) NewInstance() base.RequestInterface {
	return new(DiscoverServerRequest)
}

// IsValid return if the request is valid.
func (dto *DiscoverServerRequest) IsValid() *base.Message {
	return nil
}

// DebugInfo return the name for debug.
func (dto *DiscoverServerRequest) DebugInfo() string {
	return dto.Hostname
}
