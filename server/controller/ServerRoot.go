package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// TODO
// Should not support Post().

var (
	serverService = &service.Server{
		CRUDService: base.CRUDService{
			TemplateImpl: new(service.Server),
		},
	}
)

// ServerRoot is ther servergroup controller.
type ServerRoot struct {
}

// ResourceName returns the name this controller handle of.
func (c *ServerRoot) ResourceName() string {
	return "server"
}

// Request creates a new request DTO.
func (c *ServerRoot) Request() base.PostRequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *ServerRoot) Response() base.GetResponseInterface {
	return new(dto.GetServerResponse)
}

// Service returns the service.
func (c *ServerRoot) Service() base.CRUDServiceInterface {
	return serverService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
