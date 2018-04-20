package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerGroupRootController is ther servergroup controller.
type ServerGroupRootController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *ServerGroupRootController) GetResourceName() string {
	return "servergroup"
}

// NewRequest creates a new request DTO.
func (c *ServerGroupRootController) NewRequest() base.RequestInterface {
	request := new(dto.PostServerGroupRequest)
	request.TemplateImpl = request
	return request
}

// NewResponse creates a new response DTO.
func (c *ServerGroupRootController) NewResponse() base.ResponseInterface {
	response := new(dto.GetServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *ServerGroupRootController) GetService() base.ServiceInterface {
	return serverGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerGroupRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
