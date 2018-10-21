package errorResp

import (
	"promise/base"
)

const (
	// ErrorResponseServerServerGroupDeleteDefault ErrorResponse ID
	ErrorResponseServerServerGroupDeleteDefault = "ServerServerGroup.ErrorResponse.DeleteDefault"
)

// NewErrorResponseServerServerGroupDeleteDefault will return an error response.
func NewErrorResponseServerServerGroupDeleteDefault() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseServerServerGroupDeleteDefault
	ret.Severity = base.SeverityWarning
	ret.Description = "Delete default server-servergroup is not allowed."
	return ret
}
