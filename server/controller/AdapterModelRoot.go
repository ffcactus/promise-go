package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

var (
	adapterModelService = &base.CRUDService{
		TemplateImpl: new(service.AdapterModel),
	}
)

// AdapterModelRoot is the root controller for AdapterModel
type AdapterModelRoot struct {
}

// ResourceName returns the name this controller handle of.
func (c *AdapterModelRoot) ResourceName() string {
	return "adapterconfig"
}

// Request creates a new request DTO.
func (c *AdapterModelRoot) Request() base.PostRequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *AdapterModelRoot) Response() base.GetResponseInterface {
	return new(dto.GetAdapterModelResponse)
}

// Service returns the service.
func (c *AdapterModelRoot) Service() base.CRUDServiceInterface {
	return adapterModelService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *AdapterModelRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetAdapterModelCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
