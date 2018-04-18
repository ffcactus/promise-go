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

// UpdateActionRequestTemplateInterface is the interface that a concrete ActionRequest should have.
type UpdateActionRequestTemplateInterface interface {
	IsValid() *Message
	GetDebugName() string
	UpdateModel(ModelInterface) error
}

// UpdateActionRequestInterface is the interface that a concrete action request should implement.
type UpdateActionRequestInterface interface {
	IsValid() *Message
	GetDebugName() string
	UpdateModel(ModelInterface) error
}

// UpdateActionRequest is the implement of ActionRequestInterface.
type UpdateActionRequest struct {
	TemplateImpl UpdateActionRequestTemplateInterface `json:"-"`
}

// IsValid checks if the request is valid.
func (dto *UpdateActionRequest) IsValid() *Message {
	return dto.TemplateImpl.IsValid()
}

// GetDebugName return the name of this DTO for debug purpose.
func (dto *UpdateActionRequest) GetDebugName() string {
	return dto.TemplateImpl.GetDebugName()
}
