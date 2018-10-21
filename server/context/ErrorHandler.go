package context

import (
	"promise/base"
)

// ErrorHandlerInterface The interface of error handler.
type ErrorHandlerInterface interface {
	AppendErrorErrorResponse(errorResp base.ErrorResponse)
	ErrorErrorResponses() *[]base.ErrorResponse
}

// ErrorHandler The implementation of error handler.
type ErrorHandler struct {
	errorResps []base.ErrorResponse
}

// AppendErrorResponse Appends error response.
func (c *Base) AppendErrorResponse(errorResp base.ErrorResponse) {
	c.errorResps = append(c.errorResps, errorResp)
}

// AppendErrorResponses Appends error response.
func (c *Base) AppendErrorResponses(errorResps []base.ErrorResponse) {
	for i := range errorResps {
		c.errorResps = append(c.errorResps, errorResps[i])
	}
}

// ErrorResponses Gets error response.
func (c *Base) ErrorResponses() []base.ErrorResponse {
	return c.errorResps
}
