package message

import (
	"promise/base"
)

const (
	// MessageAuthSuccess Success
	MessageAuthSuccess = "Auth.Message.Success"
	// MessageAuthInternalError Internel error
	MessageAuthInternalError = "Auth.Message.InternalError"
	// MessageAuthBadRequest Bad request
	MessageAuthBadRequest = "Auth.Message.BadRequest"
	// MessageAuthIncorrectCredential Incorrect credential
	MessageAuthIncorrectCredential = "Auth.Message.IncorrectCredential"
	// MessageAuthNotFoundSession Session not found
	MessageAuthNotFoundSession = "Auth.Message.NotFoundSession"
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

// NewMessageAuthInternalError Internel error
func NewMessageAuthInternalError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageAuthInternalError
	ret.Severity = base.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []base.Support{
		NewSupportAuthInternalError(),
	}
	return ret
}

// NewMessageAuthBadRequest Bad request
func NewMessageAuthBadRequest() *base.Message {
	m := base.NewMessage()
	m.ID = MessageAuthBadRequest
	m.Severity = base.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []base.Support{
		NewSupportAuthBadRequest(),
	}
	return m
}

// NewMessageAuthIncorrectCredential Incorrect credential
func NewMessageAuthIncorrectCredential() *base.Message {
	m := base.NewMessage()
	m.ID = MessageAuthIncorrectCredential
	m.Severity = base.SeverityNormal
	m.Description = "Incorrect credential."
	m.Supports = []base.Support{
		NewSupportAuthIncorrectCredential(),
	}
	return m
}

// NewMessageAuthNotFoundSession Session not found
func NewMessageAuthNotFoundSession() *base.Message {
	m := base.NewMessage()
	m.ID = MessageAuthNotFoundSession
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
