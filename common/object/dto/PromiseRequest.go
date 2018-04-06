package dto

import (
	"promise/common/object/message"
)

// PromiseRequest represents the common request DTO used in Promise project.
type PromiseRequest struct {
}

// PromiseRequestInterface is the interface.
type PromiseRequestInterface interface {
	Validate() *message.Message
}

// Validate the request.
func (dto *PromiseRequest) Validate() *message.Message {
	return nil
}
