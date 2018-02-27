package model

import (
	"net/http"
	commonM "promise/common/object/model"
)

const (
	// IDMessageAuthSuccess Success
	IDMessageAuthSuccess = "IDMessageAuthSuccess"
	// IDMessageAuthInternalError Internel error
	IDMessageAuthInternalError = "IDMessageAuthInternalError"
	// IDMessageAuthBadRequest Bad request
	IDMessageAuthBadRequest = "IDMessageAuthBadRequest"
	// IDMessageAuthIncorrectCredential Incorrect credential
	IDMessageAuthIncorrectCredential = "IDMessageAuthIncorrectCredential"
	// IDMessageAuthNotFoundSession Session not found
	IDMessageAuthNotFoundSession = "IDMessageAuthNotFoundSession"
)

// NewMessageAuthInternalError Internel error
func NewMessageAuthInternalError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryAuth)
	ret.ID = IDMessageAuthInternalError
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = commonM.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []commonM.Support{
		NewSupportAuthInternalError(),
	}
	return ret
}

// NewMessageAuthBadRequest Bad request
func NewMessageAuthBadRequest() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryAuth)
	m.ID = IDMessageAuthBadRequest
	m.StatusCode = http.StatusBadRequest
	m.Severity = commonM.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []commonM.Support{
		NewSupportAuthBadRequest(),
	}
	return m
}

// NewMessageAuthIncorrectCredential Incorrect credential
func NewMessageAuthIncorrectCredential() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryAuth)
	m.ID = IDMessageAuthIncorrectCredential
	m.StatusCode = http.StatusBadRequest
	m.Severity = commonM.SeverityNormal
	m.Description = "Incorrect credential."
	m.Supports = []commonM.Support{
		NewSupportAuthIncorrectCredential(),
	}
	return m
}

// NewMessageAuthNotFoundSession Session not found
func NewMessageAuthNotFoundSession() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryAuth)
	m.ID = IDMessageAuthNotFoundSession
	m.StatusCode = http.StatusBadRequest
	m.Severity = commonM.SeverityNormal
	m.Description = "Session not found."
	m.Supports = []commonM.Support{
		NewSupportAuthNotFoundSession(),
	}
	return m
}
