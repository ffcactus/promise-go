package model

import (
	"net/http"
	commonM "promise/common/object/model"
)

const (
	// MessageIDServerSuccess Message ID
	MessageIDServerSuccess = "MessageIDServerSuccess"
	// MessageIDServerInternalError Message ID
	MessageIDServerInternalError = "MessageIDServerInternalError"
	// MessageIDServerParameterError Message ID
	MessageIDServerParameterError = "MessageIDServerParameterError"
	// MessageIDServerPostFailed Message ID
	MessageIDServerPostFailed = "MessageIDServerPostFailed"
	// MessageIDServerServerExist Message ID
	MessageIDServerServerExist = "MessageIDServerServerExist"
	// MessageIDServerServerNotExist Message ID
	MessageIDServerServerNotExist = "MessageIDServerServerNotExist"
	// MessageIDServerServerLockFailed Message ID
	MessageIDServerServerLockFailed = "MessageIDServerServerLockFailed"
	// MessageIDServerServerManagementAccountExist Message ID
	MessageIDServerServerManagementAccountExist = "MessageIDServerServerManagementAccountExist"
	// MessageIDServerServerRefreshTaskFailed Message ID
	MessageIDServerServerRefreshTaskFailed = "MessageIDServerServerRefreshTaskFailed"
)

const (
	// MessageIDServerGroupExist means server group exist.
	MessageIDServerGroupExist = "MessageIDServerGroupExist"
)

// NewArgumentServerID Get argument by server.
func NewArgumentServerID(s *Server) commonM.Argument {
	return commonM.Argument{Type: "URI", Name: s.Name, Value: s.URI}
}

// NewInternalError Create internel error.
func NewInternalError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerInternalError
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = commonM.SeverityCritical
	ret.Description = "Internal error."
	ret.Supports = []commonM.Support{
		NewSupportInternalError(),
	}
	return ret
}

// NewPostFailed create new message.
func NewPostFailed() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerPostFailed
	ret.StatusCode = http.StatusInternalServerError
	ret.Severity = commonM.SeverityWarning
	ret.Description = "Post server failed."
	ret.Supports = []commonM.Support{
		NewSupportUnableConnect(),
		NewSupportUnknownProtocol(),
		NewSupportNoBasicInfo()}
	return ret
}

// NewServerExist create new message.
func NewServerExist(s *Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerServerExist
	ret.StatusCode = http.StatusOK
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server already exists. See {0}"
	ret.Arguments = []commonM.Argument{NewArgumentServerID(s)}
	return ret
}

// NewServerNotExist create new message.
func NewServerNotExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerServerNotExist
	ret.StatusCode = http.StatusNotFound
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server doesn't exist."
	return ret
}

// NewServerLockFailed create new message.
func NewServerLockFailed(s *Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerServerLockFailed
	ret.StatusCode = http.StatusConflict
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server {0} failed to lock, server state = {1}."
	ret.Arguments = []commonM.Argument{
		NewArgumentServerID(s),
		{Type: "String", Name: s.State, Value: s.State},
	}
	ret.Supports = []commonM.Support{
		NewSupportUnableConnect(),
		NewSupportUnknownProtocol(),
		NewSupportNoBasicInfo()}
	return ret
}

// NewServerManagementAccountExist create new message.
func NewServerManagementAccountExist(s *Server) commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerServerManagementAccountExist
	ret.StatusCode = http.StatusConflict
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server {0} failed to create management account."
	ret.Arguments = []commonM.Argument{
		{Type: "String", Name: s.Name, Value: s.Name},
	}
	ret.Supports = []commonM.Support{
		NewSupportManagementAccountExist1(),
		NewSupportManagementAccountExist2(),
	}
	return ret
}

// NewServerParameterError create new message.
func NewServerParameterError() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerParameterError
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Input parameter error."
	ret.Supports = []commonM.Support{
		NewSupportParameterError(),
	}
	return ret
}

// NewServerRefreshTaskFailed create new message.
func NewServerRefreshTaskFailed() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerServerRefreshTaskFailed
	ret.StatusCode = http.StatusBadRequest
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Failed to create refresh task"
	ret.Supports = []commonM.Support{
		NewSupportInternalError(),
	}
	return ret
}

func NewServerGroupExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerGroupExist
	ret.StatusCode = http.StatusOK
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group already exists."
	return ret
}
