package errorResp

import (
	"promise/base"
)

const (
	// ErrorResponseTaskNoStep is an error response ID.
	ErrorResponseTaskNoStep = "Task.ErrorResponse.NoStep"
)

// NewErrorResponseTaskNoStep returns a new error response.
func NewErrorResponseTaskNoStep() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseTaskNoStep
	ret.Severity = base.SeverityNormal
	ret.Description = "No task steps included.."
	ret.Supports = []base.Support{}
	return ret
}
