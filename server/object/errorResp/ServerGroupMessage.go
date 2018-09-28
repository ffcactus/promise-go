package errorResp

import (
	"promise/base"
)

const (
	// ErrorResponseServerGroupDeleteDefault ErrorResponse ID
	ErrorResponseServerGroupDeleteDefault = "ServerGroup.ErrorResponse.DeleteDefault"
)

// NewErrorResponseServerGroupDeleteDefault will return an error response.
func NewErrorResponseServerGroupDeleteDefault() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseServerGroupDeleteDefault
	ret.Severity = base.SeverityWarning
	ret.Description = "Delete default server group is not allowed."
	return ret
}
