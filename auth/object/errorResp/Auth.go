package errorResp

import (
	"promise/base"
)

const (
	// ErrorResponseAuthSuccess Success
	ErrorResponseAuthSuccess = "Auth.ErrorResponse.Success"
	// ErrorResponseAuthInternalError Internel error
	ErrorResponseAuthInternalError = "Auth.ErrorResponse.InternalError"
	// ErrorResponseAuthBadRequest Bad request
	ErrorResponseAuthBadRequest = "Auth.ErrorResponse.BadRequest"
	// ErrorResponseAuthIncorrectCredential Incorrect credential
	ErrorResponseAuthIncorrectCredential = "Auth.ErrorResponse.IncorrectCredential"
	// ErrorResponseAuthNotFoundSession Session not found
	ErrorResponseAuthNotFoundSession = "Auth.ErrorResponse.NotFoundSession"
)

const (
	// SupportAuthInternalError Internel error
	SupportAuthInternalError = "Auth.Support.InternalError"
	// SupportAuthBadRequest Bad request
	SupportAuthBadRequest = "Auth.Support.BadRequest"
	// SupportAuthIncorrectCredential Incorrect credential
	SupportAuthIncorrectCredential = "Auth.Support.IncorrectCredential"
	// SupportAuthNotFoundSession Session not found
	SupportAuthNotFoundSession = "Auth.Support.NotFoundSession"
)

// NewErrorResponseAuthInternalError Internel error
func NewErrorResponseAuthInternalError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseAuthInternalError
	ret.Severity = base.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []base.Support{
		NewSupportAuthInternalError(),
	}
	return ret
}

// NewErrorResponseAuthBadRequest Bad request
func NewErrorResponseAuthBadRequest() *base.ErrorResponse {
	m := base.NewErrorResponse()
	m.ID = ErrorResponseAuthBadRequest
	m.Severity = base.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []base.Support{
		NewSupportAuthBadRequest(),
	}
	return m
}

// NewErrorResponseAuthIncorrectCredential Incorrect credential
func NewErrorResponseAuthIncorrectCredential() *base.ErrorResponse {
	m := base.NewErrorResponse()
	m.ID = ErrorResponseAuthIncorrectCredential
	m.Severity = base.SeverityNormal
	m.Description = "Incorrect credential."
	m.Supports = []base.Support{
		NewSupportAuthIncorrectCredential(),
	}
	return m
}

// NewErrorResponseAuthNotFoundSession Session not found
func NewErrorResponseAuthNotFoundSession() *base.ErrorResponse {
	m := base.NewErrorResponse()
	m.ID = ErrorResponseAuthNotFoundSession
	m.Severity = base.SeverityNormal
	m.Description = "Session not found."
	m.Supports = []base.Support{
		NewSupportAuthNotFoundSession(),
	}
	return m
}

// NewSupportAuthInternalError Internel error
func NewSupportAuthInternalError() base.Support {
	ret := base.NewSupport()
	ret.ID = SupportAuthInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportAuthBadRequest Bad request
func NewSupportAuthBadRequest() base.Support {
	ret := base.NewSupport()
	ret.ID = SupportAuthBadRequest
	ret.Reason = "Invalid request."
	ret.Solution = "Check the request and correct it."
	return ret
}

// NewSupportAuthIncorrectCredential Incorrect credential
func NewSupportAuthIncorrectCredential() base.Support {
	ret := base.NewSupport()
	ret.ID = SupportAuthIncorrectCredential
	ret.Reason = "Incorrect credential."
	ret.Solution = "Correct the credential and try again."
	return ret
}

// NewSupportAuthNotFoundSession Session not found
func NewSupportAuthNotFoundSession() base.Support {
	ret := base.NewSupport()
	ret.ID = SupportAuthNotFoundSession
	ret.Reason = "Session not found."
	ret.Solution = "Provide the right key."
	return ret
}
