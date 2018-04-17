package base

import (
// log "github.com/sirupsen/logrus"
)

// RequestTemplateInterface is the interface that a concrete Request should have.
type RequestTemplateInterface interface {
	IsValid() *Message
	GetDebugName() string
	ToModel() ModelInterface
}

// RequestInterface is the interface that  Request should have.
type RequestInterface interface {
	IsValid() *Message
	GetDebugName() string
	ToModel() ModelInterface
}

// Request is the request DTO used in Promise project.
type Request struct {
	TemplateImpl RequestInterface `json:"-"`
}

// GetDebugName return the name for debug.
func (dto *Request) GetDebugName() string {
	return dto.TemplateImpl.GetDebugName()
}

// ToModel convert the DTO to model.
func (dto *Request) ToModel() ModelInterface {
	return dto.TemplateImpl.ToModel()
}

// ActionRequestTemplateInterface is the interface that a concrete ActionRequest should have.
type ActionRequestTemplateInterface interface {
	IsValid() *Message
	GetDebugName() string
	UpdateModel(ModelInterface) error
}

// ActionRequestInterface is the interface that a concrete action request should implement.
type ActionRequestInterface interface {
	IsValid() *Message
	GetDebugName() string
	UpdateModel(ModelInterface) error
}

// ActionRequest is the implement of ActionRequestInterface.
type ActionRequest struct {
	TemplateImpl ActionRequestTemplateInterface `json:"-"`
}

// IsValid checks if the request is valid.
func (dto *ActionRequest) IsValid() *Message {
	return dto.TemplateImpl.IsValid()
}

// GetDebugName return the name of this DTO for debug purpose.
func (dto *ActionRequest) GetDebugName() string {
	return dto.TemplateImpl.GetDebugName()
}
