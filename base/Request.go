package base

import (
	log "github.com/sirupsen/logrus"
)

// RequestInterface is the interface that a PromiseRequest should have.
type RequestInterface interface {
	GetDebugName() string
	ToModel() ModelInterface
}

// Request is the request DTO used in Promise project.
type Request struct {
}

// GetDebugName return the name for debug.
func (dto *Request) GetDebugName() string {
	log.Error("Using default Request.GetDebugName().")
	return "NotProvided"
}

// ToModel convert the DTO to model.
func (dto *Request) ToModel() ModelInterface {
	var m Model
	log.Error("Using default Request.ToModel().")
	return &m
}
