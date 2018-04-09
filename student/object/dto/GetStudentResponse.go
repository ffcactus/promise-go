package dto

import (
	"promise/base"
)

// GetStudentResponse is the DTO of get student response.
type GetStudentResponse struct {
	base.Response
	Name string
	Age int
}

// GetDebugName return the name for debug.
func (dto *GetStudentResponse) GetDebugName() string {
	return dto.Name
}
