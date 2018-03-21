package message

import (
	"net/http"
	commonMessage "promise/common/object/message"
	"promise/common/object/constValue"
	"promise/server/object/model"
)

const (
	// MessageServerSuccess Message ID
	MessageServerSuccess = "MessageServerSuccess"
	// MessageServerInternalError Message ID
	MessageServerInternalError = "MessageServerInternalError"
	// MessageServerParameterError Message ID
	MessageServerParameterError = "MessageServerParameterError"
	// MessageServerPostFailed Message ID
	MessageServerPostFailed = "MessageServerPostFailed"
	// MessageServerExist Message ID
	MessageServerExist = "MessageServerExist"
	// MessageServerNotExist Message ID
	MessageServerNotExist = "MessageServerNotExist"
	// MessageServerLockFailed Message ID
	MessageServerLockFailed = "LockFailed"
	// MessageServerAccountExist Message ID
	MessageServerAccountExist = "MessageServerAccountExist"
	// MessageServerRefreshTaskFailed Message ID
	MessageServerRefreshTaskFailed = "MessageServerRefreshTaskFailed"
)

// NewArgumentServerID Get argument by server.
func NewArgumentServerID(s *model.Server) commonMessage.Argument {
	return commonMessage.Argument{Type: "URI", Name: s.Name, Value: constValue.ToServerURI(s.ID)}
}

// NewServerInternalError Create internel error.
func NewServerInternalError() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerInternalError
	ret.Severity = constValue.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []commonMessage.Support{
		NewSupportServerInternalError(),
	}
	return ret
}

// NewServerPostFailed create new message.
func NewServerPostFailed() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerPostFailed
	ret.Severity = constValue.SeverityWarning
	ret.Description = "Post server failed."
	ret.Supports = []commonMessage.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo()}
	return ret
}

// NewServerExist create new message.
func NewServerExist(s *model.Server) commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerExist
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server already ServerExists. See {0}"
	ret.Arguments = []commonMessage.Argument{NewArgumentServerID(s)}
	return ret
}

// NewServerNotExist create new message.
func NewServerNotExist() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerNotExist
	ret.StatusCode = http.StatusNotFound
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server doesn't ServerExist."
	return ret
}

// NewServerLockFailed create new message.
func NewServerLockFailed(s *model.Server) commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerLockFailed
	ret.StatusCode = http.StatusConflict
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server {0} failed to lock, server state = {1}."
	ret.Arguments = []commonMessage.Argument{
		NewArgumentServerID(s),
		{Type: "String", Name: s.State, Value: s.State},
	}
	ret.Supports = []commonMessage.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo()}
	return ret
}

// NewServerAccountExist create new message.
func NewServerAccountExist(s *model.Server) commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerAccountExist
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server {0} failed to create management account."
	ret.Arguments = []commonMessage.Argument{
		{Type: "String", Name: s.Name, Value: s.Name},
	}
	ret.Supports = []commonMessage.Support{
		NewSupportServerAccountExist1(),
		NewSupportServerAccountExist2(),
	}
	return ret
}

// NewServerParameterError create new message.
func NewServerParameterError() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerParameterError
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Input parameter error."
	ret.Supports = []commonMessage.Support{
		NewSupportServerParameterError(),
	}
	return ret
}

// NewServerRefreshTaskFailed create new message.
func NewServerRefreshTaskFailed() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerRefreshTaskFailed
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Failed to create refresh task"
	ret.Supports = []commonMessage.Support{
		NewSupportServerInternalError(),
	}
	return ret
}
