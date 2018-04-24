package message

import (
	"promise/base"
)

const (
	// MessageServerServerGroupDeleteDefault Message ID
	MessageServerServerGroupDeleteDefault = "ServerServerGroup.Message.DeleteDefault"
)

// NewMessageServerServerGroupDeleteDefault will return a message.
func NewMessageServerServerGroupDeleteDefault() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageServerServerGroupDeleteDefault
	ret.Severity = base.SeverityWarning
	ret.Description = "Delete default server-servergroup is not allowed."
	return ret
}
