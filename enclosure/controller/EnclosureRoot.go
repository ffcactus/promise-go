package controller

import (
	"promise/base"
	"promise/enclosure/object/dto"
	"promise/enclosure/service"	
)

var (
	enclosureService = &service.Enclosure {
		CRUDService: base.CRUDService{
			TemplateImpl: new(service.Enclosure),
		},
	}
)

// EnclosureRoot is the root controller for enclosure.
type EnclosureRoot struct {

}

// ResourceName returns the name this controller handle of.
func (c *EnclosureRoot) ResourceName() string {
	return "enclosure"
}

// Request creates a new request DTO.
func (c *EnclosureRoot) Request() base.PostRequestInterface {
	return nil
}

// Response creates a new response DTO.
func (c *EnclosureRoot) Response() base.GetResponseInterface {
	return new(dto.GetEnclosureResponse)
}

// Service returns the service.
func (c *EnclosureRoot) Service() base.CRUDServiceInterface {
	return enclosureService
}

// ConvertCollectionModel convert data to concrete DTO.
func (c *EnclosureRoot) ConvertCollectionModel(m *base.CollectionModel) (interface{}, error) {
	ret := new(dto.GetEnclosureCollectionResponse)
	if err := ret.Load(m); err != nil {
		return nil, err
	}
	return ret, nil
}
