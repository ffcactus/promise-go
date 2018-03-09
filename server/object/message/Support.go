package message

import (
	commonM "promise/common/object/model"
)

const (
	// IDSupportServerInternalError Support ID enum.
	IDSupportServerInternalError = "IDSupportServerInternalError"
	// IDSupportServerParameterError Support ID enum.
	IDSupportServerParameterError = "IDSupportServerParameterError"
	// IDSupportServerUnableConnect Support ID enum.
	IDSupportServerUnableConnect = "IDSupportServerUnableConnect"
	// IDSupportServerUnknownProtocol Support ID enum.
	IDSupportServerUnknownProtocol = "IDSupportServerUnknownProtocol"
	// IDSupportServerNoBasicInfo Support ID enum.
	IDSupportServerNoBasicInfo = "IDSupportServerNoBasicInfo"
	// IDSupportServerWaitForReady Support ID enum.
	IDSupportServerWaitForReady = "IDSupportServerWaitForReady"
	// IDSupportServerAccountExist1 Support ID enum.
	IDSupportServerAccountExist1 = "IDSupportServerAccountExist1"
	// IDSupportServerAccountExist2 Support ID enum.
	IDSupportServerAccountExist2 = "IDSupportServerAccountExist2"
)

// NewSupportServerInternalError Create a new support.
func NewSupportServerInternalError() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerInternalError
	ret.Reason = "An internal error happened."
	ret.Solution = "Contact Support."
	return ret
}

// NewSupportServerParameterError Create a new support.
func NewSupportServerParameterError() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerParameterError
	ret.Reason = "Input parameter may missing, incorrect or unrecognized."
	ret.Solution = "Provide the right parameter."
	return ret
}

// NewSupportServerUnableConnect Create a new support.
func NewSupportServerUnableConnect() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerUnableConnect
	ret.Reason = "Unable to connect."
	ret.Solution = "Make sure the connection is OK."
	return ret
}

// NewSupportServerUnknownProtocol Create a new support.
func NewSupportServerUnknownProtocol() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerUnknownProtocol
	ret.Reason = "Unknonw protocol."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerNoBasicInfo Create a new support.
func NewSupportServerNoBasicInfo() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerNoBasicInfo
	ret.Reason = "Failed to get basic information."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerWaitForReady Create a new support.
func NewSupportServerWaitForReady() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerWaitForReady
	ret.Reason = "The server is locked for operation."
	ret.Solution = "Wait till the operation is done."
	return ret
}

// NewSupportServerAccountExist1 Create a new support.
func NewSupportServerAccountExist1() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerAccountExist1
	ret.Reason = "The server is managed by another Director."
	ret.Solution = "Remove the server from other Director."
	return ret
}

// NewSupportServerAccountExist2 Create a new support.
func NewSupportServerAccountExist2() commonM.Support {
	ret := commonM.Support{}
	ret.ID = IDSupportServerAccountExist2
	ret.Reason = "A same management account created on the server."
	ret.Solution = "Remove the account manually or contact support."
	return ret
}
