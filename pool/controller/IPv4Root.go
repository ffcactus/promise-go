package controller

import (
	"promise/base"
	"promise/pool/object/dto"
	"promise/pool/service"
)

var (
	// ipv4Service is the service used in Student controller.
	ipv4Service = &base.CRUDService{
		TemplateImpl: new(service.IPv4Pool),
	}
)

// IPv4Root is the ipv4 pool controller.
type IPv4Root struct {
}

// ResourceName returns the name this controller handle of.
func (c *IPv4Root) ResourceName() string {
	return "ipv4"
}

// Request creates a new request DTO.
func (c *IPv4Root) Request() base.PostRequestInterface {
	return new(dto.PostIPv4PoolRequest)
}

// Response creates a new response DTO.
func (c *IPv4Root) Response() base.GetResponseInterface {
	return new(dto.GetIPv4PoolResponse)
}

// Service returns the service.
func (c *IPv4Root) Service() base.CRUDServiceInterface {
	return ipv4Service
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *IPv4Root) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetIPv4PoolCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
