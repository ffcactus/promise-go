package controller

import (
	"promise/base"
	"promise/director/object/dto"
	"promise/director/service"
)

var (
	nodeService = &service.Node{
		CRUDService: base.CRUDService{
			TemplateImpl: new(service.Node),
		},
	}
)

// NodeRoot is the root controller for Node.
type NodeRoot struct {
	base.RootController
}

// ResourceName returns the name this controller handle of.
func (c *NodeRoot) ResourceName() string {
	return "director"
}

// Request creates a new request DTO.
// We don't support post Node directly.
func (c *NodeRoot) Request() base.PostRequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *NodeRoot) Response() base.GetResponseInterface {
	return new(dto.GetNodeResponse)
}

// Service returns the service.
func (c *NodeRoot) Service() base.CRUDServiceInterface {
	return nodeService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *NodeRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetNodeCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}

// Post override the default behavior
func (c *NodeRoot) Post() {
	errorResps := []base.ErrorResponse{*base.NewErrorResponseMethodNotAllowed()}
	c.Data["json"] = &errorResps
	c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
	c.ServeJSON()
	return
}

// Delete override the default behavior
func (c *NodeRoot) Delete() {
	errorResps := []base.ErrorResponse{*base.NewErrorResponseMethodNotAllowed()}
	c.Data["json"] = &errorResps
	c.Ctx.Output.SetStatus(errorResps[0].StatusCode)
	c.ServeJSON()
	return
}
