package message

import (
	"promise/base"
)

const (
	// MessageServerGroupDeleteDefault Message ID
	MessageServerGroupDeleteDefault = "ServerGroup.Message.DeleteDefault"
)

// NewMessageServerGroupDeleteDefault will return a message.
func NewMessageServerGroupDeleteDefault() base.Message {
	ret := base.NewMessage()
	ret.ID = MessageServerGroupDeleteDefault
	ret.Severity = base.SeverityWarning
	ret.Description = "Delete default server group is not allowed."
	return ret
}
