package message

import (
	"net/http"
	commonM "promise/common/object/model"
	"promise/server/object/model"
)

const (
	// IDMessageServerSuccess Message ID
	IDMessageServerSuccess = "IDMessageServerSuccess"
	// IDMessageServerInternalError Message ID
	IDMessageServerInternalError = "IDMessageServerInternalError"
	// IDMessageServerParameterError Message ID
	IDMessageServerParameterError = "IDMessageServerParameterError"
	// IDMessageServerPostFailed Message ID
	IDMessageServerPostFailed = "IDMessageServerPostFailed"
	// IDMessageServerExist Message ID
	IDMessageServerExist = "IDMessageServerExist"
	// IDMessageServerNotExist Message ID
	IDMessageServerNotExist = "IDMessageServerNotExist"
	// IDMessageServerLockFailed Message ID
	IDMessageServerLockFailed = "LockFailed"
	// IDMessageServerAccountExist Message ID
	IDMessageServerAccountExist = "IDMessageServerAccountExist"
	// IDMessageServerRefreshTaskFailed Message ID
	IDMessageServerRefreshTaskFailed = "IDMessageServerRefreshTaskFailed"
)

const (
	// MessageIDServerGroupServerExist means server group ServerExist.
	MessageIDServerGroupServerExist = "MessageIDServerGroupServerExist"
)

// NewArgumentServerID Get argument by server.
func NewArgumentServerID(s *model.Server) commonM.Argument {
	return commonM.Argument{Type: "URI", Name: s.Name, Value: s.URI}
}

// NewServerInternalError Create internel error.
func NewServerInternalError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerInternalError
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []commonM.Support{
		NewSupportServerInternalError(),
	}
	return ret
}

// NewServerPostFailed create new message.
func NewServerPostFailed() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerPostFailed
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityWarning
	ret.Description = "Post server failed."
	ret.Supports = []commonM.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo()}
	return ret
}

// NewServerExist create new message.
func NewServerExist(s *model.Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerExist
	ret.StatusCode = http.StatusOK
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server already ServerExists. See {0}"
	ret.Arguments = []commonM.Argument{NewArgumentServerID(s)}
	return ret
}

// NewServerNotExist create new message.
func NewServerNotExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerNotExist
	ret.StatusCode = http.StatusNotFound
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server doesn't ServerExist."
	return ret
}

// NewServerLockFailed create new message.
func NewServerLockFailed(s *model.Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerLockFailed
	ret.StatusCode = http.StatusConflict
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server {0} failed to lock, server state = {1}."
	ret.Arguments = []commonM.Argument{
		NewArgumentServerID(s),
		{Type: "String", Name: s.State, Value: s.State},
	}
	ret.Supports = []commonM.Support{
		NewSupportServerUnableConnect(),
		NewSupportServerUnknownProtocol(),
		NewSupportServerNoBasicInfo()}
	return ret
}

// NewServerAccountExist create new message.
func NewServerAccountExist(s *model.Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerAccountExist
	ret.StatusCode = http.StatusConflict
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server {0} failed to create management account."
	ret.Arguments = []commonM.Argument{
		{Type: "String", Name: s.Name, Value: s.Name},
	}
	ret.Supports = []commonM.Support{
		NewSupportServerAccountExist1(),
		NewSupportServerAccountExist2(),
	}
	return ret
}

// NewServerParameterError create new message.
func NewServerParameterError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerParameterError
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Input parameter error."
	ret.Supports = []commonM.Support{
		NewSupportServerParameterError(),
	}
	return ret
}

// NewServerRefreshTaskFailed create new message.
func NewServerRefreshTaskFailed() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = IDMessageServerRefreshTaskFailed
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Failed to create refresh task"
	ret.Supports = []commonM.Support{
		NewSupportServerInternalError(),
	}
	return ret
}
