package model

import (
	commonM "promise/common/object/model"
)

const (
	// IDSupportTaskInternalError Support ID enum.
	IDSupportTaskInternalError = "IDSupportTaskInternalError"
	// IDSupportTaskSaveFailure Support ID enum.
	IDSupportTaskSaveFailure = "IDSupportTaskSaveFailure"
	// IDSupportTaskBadRequest Support ID enum.
	IDSupportTaskBadRequest = "IDSupportTaskSaveFailure"
	// IDSupportTaskNotExist Support ID enum.
	IDSupportTaskNotExist = "IDSupportTaskSaveFailure"
	// IDSupportTaskRequestMissingProperty Support ID enum.
	IDSupportTaskRequestMissingProperty = "IDSupportTaskRequestMissingProperty"
)

// NewSupportTaskInternalError Create new support.
func NewSupportTaskInternalError() commonM.Support {
	ret := commonM.NewSupport()
	ret.ID = IDSupportTaskInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportTaskSaveFailure Create new support.
func NewSupportTaskSaveFailure() commonM.Support {
	ret := commonM.NewSupport()
	ret.ID = IDSupportTaskSaveFailure
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportTaskNotExist Create new support.
func NewSupportTaskNotExist() commonM.Support {
	ret := commonM.NewSupport()
	ret.ID = IDSupportTaskNotExist
	ret.Reason = "Task ID not exist."
	ret.Solution = "Provide the right task ID."
	return ret
}

// NewSupportTaskBadRequest Create new support.
func NewSupportTaskBadRequest() commonM.Support {
	ret := commonM.NewSupport()
	ret.ID = IDSupportTaskBadRequest
	ret.Reason = "Invalid request."
	ret.Solution = "Check the request and correct it."
	return ret
}

// NewSupportTaskRequestMissingProperty Create new support.
func NewSupportTaskRequestMissingProperty() commonM.Support {
	ret := commonM.NewSupport()
	ret.ID = IDSupportTaskRequestMissingProperty
	ret.Reason = "Missing property in the request."
	ret.Solution = "Check the request and correct it."
	return ret
}
