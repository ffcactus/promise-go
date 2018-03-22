package message

import (
	"promise/common/object/constValue"
	commonMessage "promise/common/object/message"
)

const (
	// MessageServerGroupDeleteDefault Message ID
	MessageServerGroupDeleteDefault = "Server.Message.ServerGroupDeleteDefault"
)

// NewDeleteDefaultServerGroup will return a message.
func NewDeleteDefaultServerGroup() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerGroupDeleteDefault
	ret.Severity = constValue.SeverityWarning
	ret.Description = "Delete default server group is not allowed."
	return ret	
}
