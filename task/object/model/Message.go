package model

import (
	"net/http"
	commonM "promise/common/object/model"
)

const (
	// IDMessageSuccess Task message enum.
	IDMessageSuccess = "IDMessageSuccess"
	// IDMessageInternalError Task message enum.
	IDMessageInternalError = "IDMessageInternalError"
	// IDMessageTaskBadRequest Task message enum.
	IDMessageTaskBadRequest = "IDMessageTaskBadRequest"
	// IDMessageTaskSaveFailure Task message enum.
	IDMessageTaskSaveFailure = "IDMessageTaskSaveFailure"
	// IDMessageTaskNotExist Task message enum.
	IDMessageTaskNotExist = "IDMessageTaskNotExist"
	// IDMessageTaskMissingProperty Task message enum.
	IDMessageTaskMissingProperty = "IDMessageTaskRequestMissingProperty"
)

// NewMessageTaskInternalError Create new message.
func NewMessageTaskInternalError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryTask)
	ret.ID = IDMessageInternalError
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = commonM.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []commonM.Support{
		NewSupportTaskInternalError(),
	}
	return ret
}

// NewMessageTaskBadRequest Create new message.
func NewMessageTaskBadRequest() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryTask)
	m.ID = IDMessageTaskBadRequest
	m.StatusCode = http.StatusBadRequest
	m.Severity = commonM.SeverityNormal
	m.Description = "Bad request."
	m.Supports = []commonM.Support{
		NewSupportTaskBadRequest(),
	}
	return m
}

// NewMessageTaskSaveFailure Create new message.
func NewMessageTaskSaveFailure() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryTask)
	m.ID = IDMessageTaskNotExist
	m.StatusCode = http.StatusInternalServerError
	m.Severity = commonM.SeverityCritical
	m.Description = "Failed to save task."
	m.Supports = []commonM.Support{
		NewSupportTaskSaveFailure(),
	}
	return m
}

// NewMessageTaskNotExist Create new message.
func NewMessageTaskNotExist() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryTask)
	m.ID = IDMessageTaskSaveFailure
	m.StatusCode = http.StatusNotFound
	m.Severity = commonM.SeverityNormal
	m.Description = "Task not exist."
	m.Supports = []commonM.Support{
		NewSupportTaskNotExist(),
	}
	return m
}

// NewMessageTaskMissingProperty Create new message.
func NewMessageTaskMissingProperty() commonM.Message {
	m := commonM.NewMessage(commonM.CategoryTask)
	m.ID = IDMessageTaskMissingProperty
	m.StatusCode = http.StatusNotFound
	m.Severity = commonM.SeverityNormal
	m.Description = "Missing prooperty in the request"
	m.Supports = []commonM.Support{
		NewSupportTaskNotExist(),
	}
	return m
}
