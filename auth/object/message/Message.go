package message

import (
	"net/http"
	"promise/common/object/constValue"
	"promise/common/object/message"
)

const (
	// MessageAuthSuccess Success
	MessageAuthSuccess = "MessageAuthSuccess"
	// MessageAuthInternalError Internel error
	MessageAuthInternalError = "MessageAuthInternalError"
	// MessageAuthBadRequest Bad request
	MessageAuthBadRequest = "MessageAuthBadRequest"
	// MessageAuthIncorrectCredential Incorrect credential
	MessageAuthIncorrectCredential = "MessageAuthIncorrectCredential"
	// MessageAuthNotFoundSession Session not found
	MessageAuthNotFoundSession = "MessageAuthNotFoundSession"
)

// NewMessageAuthInternalError Internel error
func NewMessageAuthInternalError() message.Message {
	ret := message.NewMessage(constValue.CategoryAA)
	ret.ID = MessageAuthInternalError
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = constValue.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []message.Support{
		NewSupportAuthInternalError(),
	}
	return ret
}

// NewMessageAuthBadRequest Bad request
func NewMessageAuthBadRequest() message.Message {
	m := message.NewMessage(constValue.CategoryAA)
	m.ID = MessageAuthBadRequest
	m.StatusCode = http.StatusBadRequest
	m.Severity = constValue.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []message.Support{
		NewSupportAuthBadRequest(),
	}
	return m
}

// NewMessageAuthIncorrectCredential Incorrect credential
func NewMessageAuthIncorrectCredential() message.Message {
	m := message.NewMessage(constValue.CategoryAA)
	m.ID = MessageAuthIncorrectCredential
	m.StatusCode = http.StatusBadRequest
	m.Severity = constValue.SeverityNormal
	m.Description = "Incorrect credential."
	m.Supports = []message.Support{
		NewSupportAuthIncorrectCredential(),
	}
	return m
}

// NewMessageAuthNotFoundSession Session not found
func NewMessageAuthNotFoundSession() message.Message {
	m := message.NewMessage(constValue.CategoryAA)
	m.ID = MessageAuthNotFoundSession
	m.StatusCode = http.StatusBadRequest
	m.Severity = constValue.SeverityNormal
	m.Description = "Session not found."
	m.Supports = []message.Support{
		NewSupportAuthNotFoundSession(),
	}
	return m
}
