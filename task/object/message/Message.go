package message

import (
	"net/http"
	"promise/common/category"
	"promise/common/object/constValue"
	"promise/common/object/message"
)

const (
	// MessageSuccess Task message enum.
	MessageSuccess = "MessageSuccess"
	// MessageInternalError Task message enum.
	MessageInternalError = "MessageInternalError"
	// MessageTaskBadRequest Task message enum.
	MessageTaskBadRequest = "MessageTaskBadRequest"
	// MessageTaskSaveFailure Task message enum.
	MessageTaskSaveFailure = "MessageTaskSaveFailure"
	// MessageTaskNotExist Task message enum.
	MessageTaskNotExist = "MessageTaskNotExist"
	// MessageTaskMissingProperty Task message enum.
	MessageTaskMissingProperty = "MessageTaskRequestMissingProperty"
)

// NewMessageTaskInternalError Create new message.
func NewMessageTaskInternalError() message.Message {
	ret := message.NewMessage(category.Task)
	ret.ID = MessageInternalError
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = constValue.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []message.Support{
		NewSupportTaskInternalError(),
	}
	return ret
}

// NewMessageTaskBadRequest Create new message.
func NewMessageTaskBadRequest() message.Message {
	m := message.NewMessage(category.Task)
	m.ID = MessageTaskBadRequest
	m.StatusCode = http.StatusBadRequest
	m.Severity = constValue.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []message.Support{
		NewSupportTaskBadRequest(),
	}
	return m
}

// NewMessageTaskSaveFailure Create new message.
func NewMessageTaskSaveFailure() message.Message {
	m := message.NewMessage(category.Task)
	m.ID = MessageTaskNotExist
	m.StatusCode = http.StatusInternalServerError
	m.Severity = constValue.SeverityCritical
	m.Description = "Failed to save task."
	m.Supports = []message.Support{
		NewSupportTaskSaveFailure(),
	}
	return m
}

// NewMessageTaskNotExist Create new message.
func NewMessageTaskNotExist() message.Message {
	m := message.NewMessage(category.Task)
	m.ID = MessageTaskSaveFailure
	m.StatusCode = http.StatusNotFound
	m.Severity = constValue.SeverityNormal
	m.Description = "Task not exist."
	m.Supports = []message.Support{
		NewSupportTaskNotExist(),
	}
	return m
}

// NewMessageTaskMissingProperty Create new message.
func NewMessageTaskMissingProperty() message.Message {
	m := message.NewMessage(category.Task)
	m.ID = MessageTaskMissingProperty
	m.StatusCode = http.StatusNotFound
	m.Severity = constValue.SeverityNormal
	m.Description = "Missing prooperty in the request"
	m.Supports = []message.Support{
		NewSupportTaskNotExist(),
	}
	return m
}
