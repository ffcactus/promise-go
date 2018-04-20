package controller

import (
	"promise/base"
	"promise/server/object/dto"
	"promise/server/service"
)

// ServerServerGroupRootController is ther servergroup controller.
type ServerServerGroupRootController struct {
}

// GetResourceName returns the name this controller handle of.
func (c *ServerServerGroupRootController) GetResourceName() string {
	return "server-servergroup"
}

// NewRequest creates a new request DTO.
func (c *ServerServerGroupRootController) NewRequest() base.RequestInterface {
	request := new(dto.PostServerServerGroupRequest)
	request.TemplateImpl = request
	return request
}

// NewResponse creates a new response DTO.
func (c *ServerServerGroupRootController) NewResponse() base.ResponseInterface {
	response := new(dto.GetServerServerGroupResponse)
	response.TemplateImpl = response
	return response
}

// GetService returns the service.
func (c *ServerServerGroupRootController) GetService() base.ServiceInterface {
	return serverServerGroupService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *ServerServerGroupRootController) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetServerServerGroupCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
