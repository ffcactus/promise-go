package context

import (
	"promise/enclosure/object/dto"
)

// Discover is the context used in discover strategy.
type Discover interface {
	Base
	GetRequest() *dto.DiscoverEnclosureRequest
}

// DiscoverImpl implements Refresh context.
type DiscoverImpl struct {
	BaseImpl
	Request         *dto.DiscoverEnclosureRequest
}

// NewDiscover creates a Discover context.
func NewDiscover(id string, request *dto.DiscoverEnclosureRequest) Discover {
	ret := DiscoverImpl{}
	ret.ID = id
	ret.Request = request
	return &ret
}

// GetRequest returns request.
func (c *DiscoverImpl) GetRequest() *dto.DiscoverEnclosureRequest {
	return c.Request
}
