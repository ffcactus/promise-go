package message

import (
	commonMessage "promise/common/object/message"
	"promise/common/object/constValue"
)

const (
	// MessageServerGroupExist is message ID.
	MessageServerGroupExist = "MessageServerGroupExist"
	// MessageServerGroupNotExist is message ID.
	MessageServerGroupNotExist = "MessageServerGroupNotExist"
)

// NewServerGroupExist return a new message.
func NewServerGroupExist() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerGroupExist
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server group already exists."
	return ret
}

// NewServerGroupNotExist return a new message.
func NewServerGroupNotExist() commonMessage.Message {
	ret := commonMessage.NewMessage(constValue.CategoryServer)
	ret.ID = MessageServerGroupNotExist
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Server group not exists."
	return ret
}
