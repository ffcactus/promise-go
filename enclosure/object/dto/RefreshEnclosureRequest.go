package dto

import (
	"fmt"
	"promise/base"
)

// RefreshEnclosureRequest is the request DTO for refreshing enclosure.
type RefreshEnclosureRequest struct {
	Targets  []string
	Username string
	Password string
}

// NewInstance returns a new instance.
func (RefreshEnclosureRequest) NewInstance() base.RequestInterface {
	return new(RefreshEnclosureRequest)
}

// IsValid return if the request is valid.
func (dto *RefreshEnclosureRequest) IsValid() *base.ErrorResponse {
	if (dto.Username != "" && dto.Password == "") || (dto.Username == "" && dto.Password != "") {
		return base.NewErrorResponseInvalidRequest()
	}
	return nil
}

// String return the name for debug.
func (dto RefreshEnclosureRequest) String() string {
	return fmt.Sprintf("%v", dto.Targets)
}
