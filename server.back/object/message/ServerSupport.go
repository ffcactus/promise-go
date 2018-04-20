package message

import (
	commonMessage "promise/common/object/message"
)

const (
	// SupportServerUnableConnect Support ID enum.
	SupportServerUnableConnect = "Server.Support.ServerUnableConnect"
	// SupportServerUnknownProtocol Support ID enum.
	SupportServerUnknownProtocol = "Server.Support.ServerUnknownProtocol"
	// SupportServerNoBasicInfo Support ID enum.
	SupportServerNoBasicInfo = "Server.Support.ServerNoBasicInfo"
	// SupportServerWaitForReady Support ID enum.
	SupportServerWaitForReady = "Server.Support.ServerWaitForReady"
	// SupportServerAccountExist1 Support ID enum.
	SupportServerAccountExist1 = "Server.Support.ServerAccountExist1"
	// SupportServerAccountExist2 Support ID enum.
	SupportServerAccountExist2 = "Server.Support.ServerAccountExist2"
)

// NewSupportServerUnableConnect Create a new support.
func NewSupportServerUnableConnect() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerUnableConnect
	ret.Reason = "Unable to connect."
	ret.Solution = "Make sure the connection is OK."
	return ret
}

// NewSupportServerUnknownProtocol Create a new support.
func NewSupportServerUnknownProtocol() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerUnknownProtocol
	ret.Reason = "Unknonw protocol."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerNoBasicInfo Create a new support.
func NewSupportServerNoBasicInfo() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerNoBasicInfo
	ret.Reason = "Failed to get basic information."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerWaitForReady Create a new support.
func NewSupportServerWaitForReady() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerWaitForReady
	ret.Reason = "The server is locked for operation."
	ret.Solution = "Wait till the operation is done."
	return ret
}

// NewSupportServerAccountExist1 Create a new support.
func NewSupportServerAccountExist1() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerAccountExist1
	ret.Reason = "The server is managed by another Director."
	ret.Solution = "Remove the server from other Director."
	return ret
}

// NewSupportServerAccountExist2 Create a new support.
func NewSupportServerAccountExist2() commonMessage.Support {
	ret := commonMessage.Support{}
	ret.ID = SupportServerAccountExist2
	ret.Reason = "A same management account created on the server."
	ret.Solution = "Remove the account manually or contact support."
	return ret
}
