package dto

import (
	"promise/base"
	"promise/server/object/model"
)

// PostAdapterConfigRequest is the DTO.
type PostAdapterConfigRequest struct {
	Name string
}

// NewInstance creates a new instance.
func (dto *PostAdapterConfigRequest) NewInstance() base.RequestInterface {
	return new(PostAdapterConfigRequest)
}

// IsValid return if the request is valid.
func (dto *PostAdapterConfigRequest) IsValid() *base.Message {
	return nil
}

// String return the name for debug.
func (dto PostAdapterConfigRequest) String() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto *PostAdapterConfigRequest) ToModel() base.ModelInterface {
	ret := model.AdapterConfig{}
	ret.Category = base.CategoryAdapterConfig
	ret.Name = dto.Name
	return &ret
}
