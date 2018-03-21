package message

import (
	"promise/common/object/message"
)

const (
	// SupportAuthInternalError Internel error
	SupportAuthInternalError = "SupportAuthInternalError"
	// SupportAuthBadRequest Bad request
	SupportAuthBadRequest = "SupportAuthBadRequest"
	// SupportAuthIncorrectCredential Incorrect credential
	SupportAuthIncorrectCredential = "SupportAuthIncorrectCredential"
	// SupportAuthNotFoundSession Session not found
	SupportAuthNotFoundSession = "SupportAuthNotFoundSession"
)

// NewSupportAuthInternalError Internel error
func NewSupportAuthInternalError() message.Support {
	ret := message.NewSupport()
	ret.ID = SupportAuthInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportAuthBadRequest Bad request
func NewSupportAuthBadRequest() message.Support {
	ret := message.NewSupport()
	ret.ID = SupportAuthBadRequest
	ret.Reason = "Invalid request."
	ret.Solution = "Check the request and correct it."
	return ret
}

// NewSupportAuthIncorrectCredential Incorrect credential
func NewSupportAuthIncorrectCredential() message.Support {
	ret := message.NewSupport()
	ret.ID = SupportAuthIncorrectCredential
	ret.Reason = "Incorrect credential."
	ret.Solution = "Correct the credential and try again."
	return ret
}

// NewSupportAuthNotFoundSession Session not found
func NewSupportAuthNotFoundSession() message.Support {
	ret := message.NewSupport()
	ret.ID = SupportAuthNotFoundSession
	ret.Reason = "Session not found."
	ret.Solution = "Provide the right key."
	return ret
}
