package message

import (
	"net/http"
	"promise/common/category"
	"promise/common/object/constValue"
	commonMessage "promise/common/object/message"
	"promise/server/object/model"
)

const (
	// MessageServerSuccess Message ID
	MessageServerSuccess = "Server.Message.ServerSuccess"
	// MessageServerPostFailed Message ID
	MessageServerPostFailed = "Server.Message.ServerPostFailed"
	// MessageServerLockFailed Message ID
	MessageServerLockFailed = "Server.Message.ServerLockFailed"
	// MessageServerAccountExist Message ID
	MessageServerAccountExist = "Server.Message.ServerAccountExist"
	// MessageServerRefreshTaskFailed Message ID
	MessageServerRefreshTaskFailed = "Server.Message.ServerRefreshTaskFailed"
)

// NewArgumentServerID Get argument by server.
func NewArgumentServerID(s *model.Server) commonMessage.Argument {
	return commonMessage.Argument{Type: "URI", Name: s.Name, Value: constValue.ToServerURI(s.ID)}
}

// NewServerPostFailed create new message.
func NewServerPostFailed() commonMessage.Message {
	ret := commonMessage.NewMessage(category.Server)
	ret.ID = MessageServerPostFailed
	ret.Severity = constValue.SeverityWarning
	ret.Description = "Post server failed."
	ret.Supports = []commonMessage.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo()}
	return ret
}

// NewServerLockFailed create new message.
func NewServerLockFailed(s *model.Server) commonMessage.Message {
	ret := commonMessage.NewMessage(category.Server)
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
	ret := commonMessage.NewMessage(category.Server)
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

// NewServerRefreshTaskFailed create new message.
func NewServerRefreshTaskFailed() commonMessage.Message {
	ret := commonMessage.NewMessage(category.Server)
	ret.ID = MessageServerRefreshTaskFailed
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Failed to create refresh task"
	ret.Supports = []commonMessage.Support{
		commonMessage.NewSupportInternalError(),
	}
	return ret
}
