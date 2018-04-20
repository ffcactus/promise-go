package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerRootController is ther servergroup controller.
type ServerRootController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *ServerRootController) GetResourceName() string {
	return "server-servergroup"
}

// NewRequest creates a new request DTO.
func (c *ServerRootController) NewRequest() base.RequestInterface {
	request := new(dto.PostServerRequest)
	request.TemplateImpl = request
	return request
}

// NewResponse creates a new response DTO.
func (c *ServerRootController) NewResponse() base.ResponseInterface {
	response := new(dto.GetServerResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *ServerRootController) GetService() base.ServiceInterface {
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
