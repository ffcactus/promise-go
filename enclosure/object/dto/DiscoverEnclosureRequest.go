package dto

import (
	"promise/base"
	"promise/enclosure/object/model"
)

// DiscoverEnclosureRequest is the DTO.
type DiscoverEnclosureRequest struct {
	Name string
	Description string
	Type string
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

// NewEnclosure creates a enclosure model based on the discover request.
func (dto DiscoverEnclosureRequest) NewEnclosure() *model.Enclosure {
	enclosure := model.Enclosure{}
	enclosure.Name = dto.Name
	enclosure.Description = dto.Description
	enclosure.Type = dto.Type
	enclosure.Addresses = []string{dto.Address}
	enclosure.Credential.Username = dto.Username
	enclosure.Credential.Password = dto.Password
	return &enclosure
}