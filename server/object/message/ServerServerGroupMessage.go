package message

import (
	"promise/common/object/constValue"
	commonMessage "promise/common/object/message"
)

const (
	// MessageServerServerGroupDeleteDefault Message ID
	MessageServerServerGroupDeleteDefault = "Server.Message.ServerServerGroupDeleteDefault"
)

// NewDeleteDefaultServerServerGroup will return a message.
func NewDeleteDefaultServerServerGroup() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerServerGroupDeleteDefault
	ret.Severity = constValue.SeverityWarning
	ret.Description = "Delete default server-servergroup is not allowed."
	return ret
}
