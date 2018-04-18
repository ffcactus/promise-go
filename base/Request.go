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
}

// ActionRequestInterface is the interface that a UpdateAction should implement.
type ActionRequestInterface interface {
	IsValid() *Message
	GetDebugName() string
}

// ActionRequest is the implement of ActionRequestInterface.
// The UpdateActionRequest is a kind of action that can be taken a parameter to
// perform an action to resource.
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

// UpdateActionRequestTemplateInterface is the interface that a concrete UpdateActionRequest should have.
type UpdateActionRequestTemplateInterface interface {
	ActionRequestTemplateInterface
	UpdateModel(ModelInterface) error
}

// UpdateActionRequestInterface is the interface that a UpdateActionRequest should implement.
type UpdateActionRequestInterface interface {
	ActionRequestInterface
	UpdateModel(ModelInterface) error
}

// UpdateActionRequest is the implement of UpdateActionRequestInterface.
// The UpdateActionRequest is a kind of action that just update the resource.
type UpdateActionRequest struct {
	TemplateImpl UpdateActionRequestTemplateInterface `json:"-"`
}