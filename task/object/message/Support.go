package message

import (
	commonMessage "promise/common/object/message"
)

const (
	// SupportTaskInternalError Support ID enum.
	SupportTaskInternalError = "SupportTaskInternalError"
	// SupportTaskSaveFailure Support ID enum.
	SupportTaskSaveFailure = "SupportTaskSaveFailure"
	// SupportTaskBadRequest Support ID enum.
	SupportTaskBadRequest = "SupportTaskSaveFailure"
	// SupportTaskNotExist Support ID enum.
	SupportTaskNotExist = "SupportTaskSaveFailure"
	// SupportTaskRequestMissingProperty Support ID enum.
	SupportTaskRequestMissingProperty = "SupportTaskRequestMissingProperty"
)

// NewSupportTaskInternalError Create new support.
func NewSupportTaskInternalError() commonMessage.Support {
	ret := commonMessage.NewSupport()
	ret.ID = SupportTaskInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportTaskSaveFailure Create new support.
func NewSupportTaskSaveFailure() commonMessage.Support {
	ret := commonMessage.NewSupport()
	ret.ID = SupportTaskSaveFailure
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportTaskNotExist Create new support.
func NewSupportTaskNotExist() commonMessage.Support {
	ret := commonMessage.NewSupport()
	ret.ID = SupportTaskNotExist
	ret.Reason = "Task ID not exist."
	ret.Solution = "Provide the right task ID."
	return ret
}

// NewSupportTaskBadRequest Create new support.
func NewSupportTaskBadRequest() commonMessage.Support {
	ret := commonMessage.NewSupport()
	ret.ID = SupportTaskBadRequest
	ret.Reason = "Invalid request."
	ret.Solution = "Check the request and correct it."
	return ret
}

// NewSupportTaskRequestMissingProperty Create new support.
func NewSupportTaskRequestMissingProperty() commonMessage.Support {
	ret := commonMessage.NewSupport()
	ret.ID = SupportTaskRequestMissingProperty
	ret.Reason = "Missing property in the request."
	ret.Solution = "Check the request and correct it."
	return ret
}
