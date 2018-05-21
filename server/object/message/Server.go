package message

import (
	"promise/base"
	"promise/server/object/model"
)

const (
	// MessageServerDiscoverFailed Message ID
	MessageServerDiscoverFailed = "Server.Message.DiscoverFailed"
	// MessageServerAccountExist Message ID
	MessageServerAccountExist = "Server.Message.AccountExist"
	// MessageServerRefreshTaskFailed Message ID
	MessageServerRefreshTaskFailed = "Server.Message.RefreshTaskFailed"
)

const (
	// SupportServerUnableConnect Support ID enum.
	SupportServerUnableConnect = "Server.Support.UnableConnect"
	// SupportServerUnknownProtocol Support ID enum.
	SupportServerUnknownProtocol = "Server.Support.UnknownProtocol"
	// SupportServerNoBasicInfo Support ID enum.
	SupportServerNoBasicInfo = "Server.Support.NoBasicInfo"
	// SupportServerWaitForReady Support ID enum.
	SupportServerWaitForReady = "Server.Support.WaitForReady"
	// SupportServerAccountExist1 Support ID enum.
	SupportServerAccountExist1 = "Server.Support.AccountExist1"
	// SupportServerAccountExist2 Support ID enum.
	SupportServerAccountExist2 = "Server.Support.AccountExist2"
)

// NewArgumentServerID Get argument by server.
func NewArgumentServerID(s *model.Server) base.Argument {
	return base.Argument{Type: "URI", Name: s.Name, Value: base.ToServerURI(s.ID)}
}

// NewMessageServerDiscoverFailed create new message.
func NewMessageServerDiscoverFailed() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageServerDiscoverFailed
	ret.Severity = base.SeverityWarning
	ret.Description = "Post server failed."
	ret.Supports = []base.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo(),
	}
	return ret
}

// NewMessageServerAccountExist create new message.
func NewMessageServerAccountExist(s *model.Server) *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageServerAccountExist
	ret.Severity = base.SeverityNormal
	ret.Description = "Server {0} failed to create management account."
	ret.Arguments = []base.Argument{
		{Type: "String", Name: s.Name, Value: s.Name},
	}
	ret.Supports = []base.Support{
		NewSupportServerAccountExist1(),
		NewSupportServerAccountExist2(),
	}
	return ret
}

// NewMessageServerRefreshTaskFailed create new message.
func NewMessageServerRefreshTaskFailed() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageServerRefreshTaskFailed
	ret.Severity = base.SeverityNormal
	ret.Description = "Failed to create refresh task"
	ret.Supports = []base.Support{
		base.NewSupportInternalError(),
	}
	return ret
}

// NewSupportServerUnableConnect Create a new support.
func NewSupportServerUnableConnect() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerUnableConnect
	ret.Reason = "Unable to connect."
	ret.Solution = "Make sure the connection is OK."
	return ret
}

// NewSupportServerUnknownProtocol Create a new support.
func NewSupportServerUnknownProtocol() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerUnknownProtocol
	ret.Reason = "Unknonw protocol."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerNoBasicInfo Create a new support.
func NewSupportServerNoBasicInfo() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerNoBasicInfo
	ret.Reason = "Failed to get basic information."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportServerWaitForReady Create a new support.
func NewSupportServerWaitForReady() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerWaitForReady
	ret.Reason = "The server is locked for operation."
	ret.Solution = "Wait till the operation is done."
	return ret
}

// NewSupportServerAccountExist1 Create a new support.
func NewSupportServerAccountExist1() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerAccountExist1
	ret.Reason = "The server is managed by another Director."
	ret.Solution = "Remove the server from other Director."
	return ret
}

// NewSupportServerAccountExist2 Create a new support.
func NewSupportServerAccountExist2() base.Support {
	ret := base.Support{}
	ret.ID = SupportServerAccountExist2
	ret.Reason = "A same management account created on the server."
	ret.Solution = "Remove the account manually or contact support."
	return ret
}
