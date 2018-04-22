package base

import (
// "time"
)

// ResponseInterface is the inteface that a Response have.
type ResponseInterface interface {
	DebugInfo() string
}

// // GetDebugName return the name for debug.
// func (dto *Response) GetDebugName() string {
// 	return dto.TemplateImpl.GetDebugName()
// }

// // GetID returns ID.
// func (dto *Response) GetID() string {
// 	return dto.ID
// }

// // GetCategory returns Category.
// func (dto *Response) GetCategory() string {
// 	return dto.Category
// }

// // Load data from model.
// func (dto *Response) Load(m ModelInterface) error {
// 	return dto.TemplateImpl.Load(m)
// }
