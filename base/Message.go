package base

import (
	"net/http"
	"time"
)

const (
	// SeverityNormal is the errorResp severity level.
	SeverityNormal = "Normal"
	// SeverityWarning is the errorResp severity level.
	SeverityWarning = "Warning"
	// SeverityCritical is the errorResp severity level.
	SeverityCritical = "Critical"
)

const (
	// ErrorResponseInternalError is errorResp ID.
	ErrorResponseInternalError = "Promise.ErrorResponse.InternalError"
	// ErrorResponseNotExist is errorResp ID.
	ErrorResponseNotExist = "Promise.ErrorResponse.NotExist"
	// ErrorResponseMethodNotAllowed is errorResp ID.
	ErrorResponseMethodNotAllowed = "Promise.ErrorResponse.MethodNotAllowed"
	// ErrorResponseDuplicate is errorResp ID.
	ErrorResponseDuplicate = "Promise.ErrorResponse.Duplicate"
	// ErrorResponseInvalidRequest is errorResp ID.
	ErrorResponseInvalidRequest = "Promise.ErrorResponse.InvalidRequest"
	// ErrorResponseUnknownPropertyValue is errorResp ID.
	ErrorResponseUnknownPropertyValue = "Promise.ErrorResponse.UnknownPropertyValue"
	// ErrorResponseTimeout is errorResp ID.
	ErrorResponseTimeout = "Promise.ErrorResponse.Timeout"
	// ErrorResponseTransactionError is errorResp ID.
	ErrorResponseTransactionError = "Promise.ErrorResponse.TransactionError"
	// ErrorResponseUnknownFilterName is errorResp ID.
	ErrorResponseUnknownFilterName = "Promise.ErrorResponse.UnknownFilterName"
	// ErrorResponseBusy is errorResp ID.
	ErrorResponseBusy = "Promise.ErrorResponse.Busy"
	// ErrorResponseErrorState is errorResp ID.
	ErrorResponseErrorState = "Promise.ErrorResponse.ErrorState"
)

const (
	// SupportNotExist is Support ID.
	SupportNotExist = "Promise.Support.NotExist"
	// SupportDuplicate is Support ID.
	SupportDuplicate = "Promise.Support.Duplicate"
	// SupportInvalidRequest is Support ID.
	SupportInvalidRequest = "Promise.Support.InvalidRequest"
	// SupportInternalError is Support ID.
	SupportInternalError = "Promise.Support.InternalError"
	// SupportTimeout is Support ID.
	SupportTimeout = "Promise.Support.Timeout"
	// SupportTransactionError is Support ID.
	SupportTransactionError = "Promise.Support.TransactionError"
	// SupportBusy is Support ID.
	SupportBusy = "Promise.Support.Busy"
	// SupportErrorState is Support ID.
	SupportErrorState = "Promise.Support.ErrorState"
)

// For auth.

// Argument is used to replace a portion in a string. For example, to express
// an error happend on a server, we need replace the {0} in the string below:
// "There is an error on server {0}". The argument here can be:
// { "Type": "URI", "Name": "Server 1", "Value": "/api/v1/server/xxxxxx" }
type Argument struct {
	Type  string `json:"Type"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Support tells how to solve a problem.
type Support struct {
	ID                string     `json:"ID"` // The unique ID within a micro service.
	Reason            string     `json:"Reason"`
	ReasonArguments   []Argument `json:"ReasonArguments"`
	Solution          string     `json:"Solution"`
	SolutionArguments []Argument `json:"SolutionArguments"`
}

// ErrorResponse is the errorResp used in Promise project.
type ErrorResponse struct {
	ID          string     `json:"ID"` // The unique ID within a micro service.
	StatusCode  int        `json:"-"`  // The HTTP status code along with this errorResp.
	Severity    string     `json:"Severity"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	Description string     `json:"Description"`
	Arguments   []Argument `json:"Arguments"` // The arguments in the description.
	Supports    []Support  `json:"Supports"`
}

// NewErrorResponse create a errorResp with default value.
func NewErrorResponse() *ErrorResponse {
	ret := ErrorResponse{
		CreatedAt:  time.Now(),
		StatusCode: http.StatusBadRequest,
	}
	return &ret
}

// NewErrorResponseNotExist returns a errorResp that means the resource does not exist.
func NewErrorResponseNotExist() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseNotExist
	ret.Severity = SeverityNormal
	ret.Description = "Resource does not exist."
	ret.StatusCode = http.StatusNotFound
	ret.Supports = []Support{
		NewSupportNotExist(),
	}
	return ret
}

// NewErrorResponseMethodNotAllowed returns a errorResp that means the method is not allowed.
func NewErrorResponseMethodNotAllowed() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseMethodNotAllowed
	ret.Severity = SeverityNormal
	ret.Description = "Method not allowed."
	ret.StatusCode = http.StatusMethodNotAllowed
	ret.Supports = []Support{}
	return ret
}

// NewErrorResponseDuplicate returns a errorResp that means resource duplication happened.
func NewErrorResponseDuplicate() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseDuplicate
	ret.Severity = SeverityNormal
	ret.Description = "Resource duplicated."
	ret.Supports = []Support{
		NewSupportDuplicate(),
	}
	return ret
}

// NewErrorResponseInvalidRequest returns a errorResp that means the request is Invalid.
func NewErrorResponseInvalidRequest() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseInvalidRequest
	ret.Severity = SeverityNormal
	ret.Description = "The request is invalid."
	ret.Supports = []Support{
		NewSupportInvalidRequest(),
	}
	return ret
}

// NewErrorResponseUnknownPropertyValue returns a errorResp that means the property in the request have unknown value.
func NewErrorResponseUnknownPropertyValue() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseUnknownPropertyValue
	ret.Severity = SeverityNormal
	ret.Description = "Unknown property value."
	ret.Supports = []Support{}
	return ret
}

// NewErrorResponseInternalError returns a errorResp that means there is a internal error happened.
func NewErrorResponseInternalError() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseInternalError
	ret.Severity = SeverityNormal
	ret.Description = "Internal error happened while process the request."
	ret.Supports = []Support{
		NewSupportInternalError(),
	}
	return ret
}

// NewErrorResponseTimeout returns a errorResp that means there is a timeout happend.
func NewErrorResponseTimeout() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseTimeout
	ret.Severity = SeverityNormal
	ret.Description = "I/O operation timeout."
	ret.Supports = []Support{
		NewSupportTimeout(),
	}
	return ret
}

// NewErrorResponseTransactionError returns a errorResp that means transaction error.
func NewErrorResponseTransactionError() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseTransactionError
	ret.Severity = SeverityNormal
	ret.Description = "Transaction error."
	ret.Supports = []Support{
		NewSupportTransactionError(),
	}
	return ret
}

// NewErrorResponseUnknownFilterName returns a new errorResp.
func NewErrorResponseUnknownFilterName() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseUnknownFilterName
	ret.Severity = SeverityNormal
	ret.Description = "Unknown filter name."
	ret.Supports = []Support{}
	return ret
}

// NewErrorResponseBusy returns a new errorResp.
func NewErrorResponseBusy() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseBusy
	ret.Severity = SeverityNormal
	ret.Description = "The system is busy."
	ret.Supports = []Support{
		NewSupportBusy(),
	}
	return ret
}

// NewErrorResponseErrorState returns a new errorResp.
func NewErrorResponseErrorState() *ErrorResponse {
	ret := NewErrorResponse()
	ret.ID = ErrorResponseErrorState
	ret.Severity = SeverityNormal
	ret.Description = "The operation is failed due to resource state."
	ret.Supports = []Support{
		NewSupportErrorState(),
	}
	return ret
}

// NewSupport create a new Support.
func NewSupport() Support {
	ret := Support{}
	ret.ReasonArguments = make([]Argument, 0)
	ret.SolutionArguments = make([]Argument, 0)
	return ret
}

// NewSupportNotExist will returns a support errorResp.
func NewSupportNotExist() Support {
	ret := NewSupport()
	ret.ID = SupportNotExist
	ret.Reason = "There is no resource match the identifier."
	ret.Solution = "Verify the identifier and try again."
	return ret
}

// NewSupportDuplicate will returns a support errorResp.
func NewSupportDuplicate() Support {
	ret := NewSupport()
	ret.ID = SupportDuplicate
	ret.Reason = "The resource duplication happend inside."
	ret.Solution = "Stop create the duplicated resource."
	return ret
}

// NewSupportInvalidRequest will returns a support errorResp.
func NewSupportInvalidRequest() Support {
	ret := NewSupport()
	ret.ID = SupportInvalidRequest
	ret.Reason = "The request is invalid."
	ret.Solution = "Verify the request and try again."
	return ret
}

// NewSupportInternalError will returns a support errorResp.
func NewSupportInternalError() Support {
	ret := NewSupport()
	ret.ID = SupportInternalError
	ret.Reason = "Internal error."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportTimeout will returns a support errorResp.
func NewSupportTimeout() Support {
	ret := NewSupport()
	ret.ID = SupportTimeout
	ret.Reason = "I/O operation timeout."
	ret.Solution = "Try again later."
	return ret
}

// NewSupportTransactionError will returns a support errorResp.
func NewSupportTransactionError() Support {
	ret := NewSupport()
	ret.ID = SupportTransactionError
	ret.Reason = "DB operation failed."
	ret.Solution = "Try again later or contact support."
	return ret
}

// NewSupportBusy will returns a support errorResp.
func NewSupportBusy() Support {
	ret := NewSupport()
	ret.ID = SupportBusy
	ret.Reason = "Too many concurrent operation."
	ret.Solution = "Try again later."
	return ret
}

// NewSupportErrorState will returns a support errorResp.
func NewSupportErrorState() Support {
	ret := NewSupport()
	ret.ID = SupportErrorState
	ret.Reason = "The operation is not allowed during resource state."
	ret.Solution = "Make sure the resource is in a suitable state."
	return ret
}
