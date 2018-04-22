package base

import (
// log "github.com/sirupsen/logrus"
)

// // RequestTemplateInterface is the interface that a concrete Request should have.
// type RequestTemplateInterface interface {
// 	IsValid() *Message
// 	GetDebugName() string
// 	ToModel() ModelInterface
// }

// RequestInterface is the interface that  Request should have.
type RequestInterface interface {
	NewInstance() RequestInterface // Create a new instance of the request.
	IsValid() *Message             // Check if the request is valid.
	DebugInfo() string             // Get the debug info.
	// ToModel() ModelInterface
}

// // Request is the request DTO used in Promise project.
// type Request struct {
// 	// TemplateImpl RequestInterface `json:"-"`
// }

// // GetDebugName return the name for debug.
// func (dto *Request) GetDebugName() string {
// 	return dto.TemplateImpl.GetDebugName()
// }

// // ToModel convert the DTO to model.
// func (dto *Request) ToModel() ModelInterface {
// 	return dto.TemplateImpl.ToModel()
// }

// // IsValid checks if the request is valid.
// func (dto *ActionRequest) IsValid() *Message {
// 	return dto.TemplateImpl.IsValid()
// }
