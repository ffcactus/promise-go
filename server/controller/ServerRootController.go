package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// TODO
// Should not support Post().

var (
	serverService = &base.CRUDService{
		TemplateImpl: new(service.Server),
	}
)

// ServerRootController is ther servergroup controller.
type ServerRootController struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerRootController) ResourceName() string {
	return "server"
}

// Request creates a new request DTO.
func (c *ServerRootController) Request() base.RequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *ServerRootController) Response() base.GetResponseInterface {
	return new(dto.GetServerResponse)
}

// Service returns the service.
func (c *ServerRootController) Service() base.CRUDServiceInterface {
	return serverService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
