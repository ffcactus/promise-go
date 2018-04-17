package controller

import (
	"promise/base"
	"promise/pool/object/dto"
	"promise/pool/service"
)

var (
	// ipv4Service is the service used in Student controller.
	ipv4Service = &base.Service{
		TemplateImpl: new(service.IPv4Pool),
	}
)

// IPv4RootController is the ipv4 pool controller.
type IPv4RootController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *IPv4RootController) GetResourceName() string {
	return "ipv4"
}

// NewRequest creates a new request DTO.
func (c *IPv4RootController) NewRequest() base.RequestInterface {
	request := new(dto.PostIPv4PoolRequest)
	request.TemplateImpl = request
	return request
}

// NewResponse creates a new response DTO.
func (c *IPv4RootController) NewResponse() base.ResponseInterface {
	response := new(dto.GetIPv4PoolResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *IPv4RootController) GetService() base.ServiceInterface {
	return ipv4Service
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *IPv4RootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetIPv4PoolCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
