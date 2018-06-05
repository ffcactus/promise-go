package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

var (
	adapterConfigService = &base.CRUDService{
		TemplateImpl: new(service.AdapterConfig),
	}
)

// AdapterConfigRoot is the root controller for AdapterConfig
type AdapterConfigRoot struct {
}

// ResourceName returns the name this controller handle of.
func (c *AdapterConfigRoot) ResourceName() string {
	return "adapterconfig"
}

// Request creates a new request DTO.
func (c *AdapterConfigRoot) Request() base.PostRequestInterface {
	return new(dto.PostAdapterConfigRequest)
}

// Response creates a new response DTO.
func (c *AdapterConfigRoot) Response() base.GetResponseInterface {
	return new(dto.GetAdapterConfigResponse)
}

// Service returns the service.
func (c *AdapterConfigRoot) Service() base.CRUDServiceInterface {
	return adapterConfigService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *AdapterConfigRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetAdapterConfigCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
