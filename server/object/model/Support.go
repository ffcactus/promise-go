package model

import (
	commonM "promise/common/object/model"
)

const (
	// IDSupportInternalError Support ID enum.
	IDSupportInternalError = "IDSupportInternalError"
	// IDSupportParameterError Support ID enum.
	IDSupportParameterError = "IDSupportParameterError"
	// IDSupportUnableConnect Support ID enum.
	IDSupportUnableConnect = "IDSupportUnableConnect"
	// IDSupportUnknownProtocol Support ID enum.
	IDSupportUnknownProtocol = "IDSupportUnknownProtocol"
	// IDSupportNoBasicInfo Support ID enum.
	IDSupportNoBasicInfo = "IDSupportNoBasicInfo"
	// IDSupportWaitForReady Support ID enum.
	IDSupportWaitForReady = "IDSupportWaitForReady"
	// IDSupportManagementAccountExist1 Support ID enum.
	IDSupportManagementAccountExist1 = "IDSupportManagementAccountExist1"
	// IDSupportManagementAccountExist2 Support ID enum.
	IDSupportManagementAccountExist2 = "IDSupportManagementAccountExist2"
)

// NewSupportInternalError Create a new support.
func NewSupportInternalError() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportParameterError Create a new support.
func NewSupportParameterError() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportParameterError
	ret.Reason = "Input parameter may missing, incorrect or unrecognized."
	ret.Solution = "Provide the right parameter."
	return ret
}

// NewSupportUnableConnect Create a new support.
func NewSupportUnableConnect() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportUnableConnect
	ret.Reason = "Unable to connect."
	ret.Solution = "Make sure the connection is OK."
	return ret
}

// NewSupportUnknownProtocol Create a new support.
func NewSupportUnknownProtocol() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportUnknownProtocol
	ret.Reason = "Unknonw protocol."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportNoBasicInfo Create a new support.
func NewSupportNoBasicInfo() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportNoBasicInfo
	ret.Reason = "Failed to get basic information."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportWaitForReady Create a new support.
func NewSupportWaitForReady() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportWaitForReady
	ret.Reason = "The server is locked for operation."
	ret.Solution = "Wait till the operation is done."
	return ret
}

// NewSupportManagementAccountExist1 Create a new support.
func NewSupportManagementAccountExist1() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportManagementAccountExist1
	ret.Reason = "The server is managed by another Director."
	ret.Solution = "Remove the server from other Director."
	return ret
}

// NewSupportManagementAccountExist2 Create a new support.
func NewSupportManagementAccountExist2() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportManagementAccountExist2
	ret.Reason = "A same management account created on the server."
	ret.Solution = "Remove the account manually or contact support."
	return ret
}
