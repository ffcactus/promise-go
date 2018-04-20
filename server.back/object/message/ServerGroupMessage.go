package message

import (
	"promise/common/object/constvalue"
	commonMessage "promise/common/object/message"
)

const (
	// MessageServerGroupDeleteDefault Message ID
	MessageServerGroupDeleteDefault = "Server.Message.ServerGroupDeleteDefault"
)

// NewDeleteDefaultServerGroup will return a message.
func NewDeleteDefaultServerGroup() commonMessage.Message {
	ret := commonMessage.NewMessage()
	ret.ID = MessageServerGroupDeleteDefault
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "Delete default server group is not allowed."
	return ret
}
