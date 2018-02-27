package model

import (
	"promise/common/object/model"
)

const (
	// IDSupportAuthInternalError Internel error
	IDSupportAuthInternalError = "IDSupportAuthInternalError"
	// IDSupportAuthBadRequest Bad request
	IDSupportAuthBadRequest = "IDSupportAuthBadRequest"
	// IDSupportAuthIncorrectCredential Incorrect credential
	IDSupportAuthIncorrectCredential = "IDSupportAuthIncorrectCredential"
	// IDSupportAuthNotFoundSession Session not found
	IDSupportAuthNotFoundSession = "IDSupportAuthNotFoundSession"
)

// NewSupportAuthInternalError Internel error
func NewSupportAuthInternalError() model.Support {
	ret := model.NewSupport()
	ret.ID = IDSupportAuthInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportAuthBadRequest Bad request
func NewSupportAuthBadRequest() model.Support {
	ret := model.NewSupport()
	ret.ID = IDSupportAuthBadRequest
	ret.Reason = "Invalid request."
	ret.Solution = "Check the request and correct it."
	return ret
}

// NewSupportAuthIncorrectCredential Incorrect credential
func NewSupportAuthIncorrectCredential() model.Support {
	ret := model.NewSupport()
	ret.ID = IDSupportAuthIncorrectCredential
	ret.Reason = "Incorrect credential."
	ret.Solution = "Correct the credential and try again."
	return ret
}

// NewSupportAuthNotFoundSession Session not found
func NewSupportAuthNotFoundSession() model.Support {
	ret := model.NewSupport()
	ret.ID = IDSupportAuthNotFoundSession
	ret.Reason = "Session not found."
	ret.Solution = "Provide the right key."
	return ret
}
